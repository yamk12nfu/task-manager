package infrastructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter() *Router {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return &Router{echo: e}
}

type Router struct {
	echo *echo.Echo
}

func (r *Router) Start() error {
	return r.echo.Start(":8000")
}
