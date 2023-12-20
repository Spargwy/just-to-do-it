package api

import (
	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/google/uuid"
)

type Tasker interface {
	TasksList() ([]*models.Task, error)
	TaskByID(id uuid.UUID) (*models.Task, error)
	CreateTask() error
}
