package service

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/google/uuid"
)

func (c *ClientExecutor) TasksList(ctx context.Context, filterParams map[string][]string, filterStruct models.Task) (models.Response, error) {
	where := buildWhereConditionFromParams(filterParams)
	tasks, err := c.db.TasksList(where, filterStruct)
	if err != nil {
		return models.Response{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return models.Response{
		Status: http.StatusOK,
		Object: tasks,
	}, nil
}

func (c *ClientExecutor) TaskByID(ctx context.Context, id uuid.UUID) (models.Response, error) {
	task, err := c.db.TaskByID(id)
	if err != nil && err == sql.ErrNoRows {
		return models.Response{
			Status:  http.StatusOK,
			Message: "task not found",
		}, nil
	}

	if err != nil {
		return models.Response{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, err
	}
	return models.Response{
		Status: http.StatusOK,
		Object: *task,
	}, nil
}

func (c *ClientExecutor) CreateTask(ctx context.Context, req models.CreateTaskRequest, user models.User) (models.Response, error) {
	task := models.Task{CreaterID: user.ID}
	task.Convert(req)
	err := c.db.CreateTask(&task)
	if err != nil {
		return models.Response{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, err
	}
	return models.Response{
		Status:  http.StatusOK,
		Message: "task created",
		Object:  task.ID,
	}, nil
}
