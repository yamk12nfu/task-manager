package infrastructure

import (
	"context"
	"net/http"
	"task-manager/app/interface/controllers"
	"task-manager/app/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *Router {
	e := echo.New()

	e.Use(
		middleware.Logger(),
	)

	sqlHandler := NewSqlHander()

	v1 := e.Group("/api/v1")

	// task
	taskController := controllers.NewTaskController(sqlHandler)
	tasks := v1.Group("/tasks")
	tasks.POST("", newHandlerFunc(taskController.Create))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return &Router{echo: e, closables: []closable{sqlHandler}}
}

type Router struct {
	echo      *echo.Echo
	closables []closable
}

type closable interface {
	Close() error
}

func (r *Router) Start() {
	go func() {
		if err := r.echo.Start(":8080"); err != http.ErrServerClosed {
			r.echo.Logger.Fatal(err)
		}
	}()
}

func (r *Router) Shutdown(ctx context.Context) {
	if err := r.echo.Shutdown(ctx); err != nil {
		r.echo.Logger.Fatal(err)
	}

	r.closeAll(ctx)
}

func (r *Router) closeAll(ctx context.Context) {
	var wg utils.WaitGroup

	for _, c := range r.closables {
		c := c
		wg.Go(func() {
			select {
			case <-ctx.Done():
				r.echo.Logger.Error(ctx.Err())
			default:
				if err := c.Close(); err != nil {
					r.echo.Logger.Error(err)
				}
			}
		})
	}

	wg.Wait()
}

type ControllerFunc func(controllers.Context) error

func newHandlerFunc(f ControllerFunc) echo.HandlerFunc {
	return func(c echo.Context) error { return f(c) }
}
