package service

import (
	"context"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/google/uuid"
)

func (c *ClientExecutor) TasksList(ctx context.Context, user models.User, filterParams map[string][]string) ([]*models.Task, error) {
	where := buildWhereConditionFromParams(filterParams)
	tasks, err := c.db.TasksList(where, user.ID)
	return tasks, err
}

func (c *ClientExecutor) TaskByID(ctx context.Context, user models.User, id uuid.UUID) (*models.Task, error) {
	task, err := c.db.TaskByID(id, user.ID)
	return task, err
}

func (c *ClientExecutor) CreateTask(ctx context.Context, req models.CreateTaskRequest, user models.User) (uuid.UUID, error) {
	task := models.Task{CreaterID: user.ID}
	task.Convert(req)
	err := c.db.CreateTask(&task)
	return task.ID, err
}
