package controllers

import (
	"net/http"
	"task-manager/app/interface/database"
	"task-manager/app/interface/schemas"
	"task-manager/app/usecases"
)

type TaskController struct {
	Interactor usecases.TaskInteractor
}

func NewTaskController(sqlHandler database.SQLHandler) *TaskController {
	return &TaskController{
		Interactor: usecases.TaskInteractor{
			Transaction: sqlHandler.Transaction(),
			TaskRepository: &database.TaskRepository{
				SQLHandler: sqlHandler,
			},
			UniqueIDRepository: &database.UniqueIDRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *TaskController) Create(c Context) error {
	taskRequest := schemas.TaskSchema{}
	if err := c.Bind(&taskRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()

	err := controller.Interactor.Save(
		ctx,
		taskRequest.Name,
		taskRequest.Description,
		taskRequest.DueDate,
		taskRequest.Status,
	)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, http.StatusText(http.StatusCreated))
}

func (controller *TaskController) Show(c Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()

	task, err := controller.Interactor.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewTaskSchemaFromEntities(task))
}
