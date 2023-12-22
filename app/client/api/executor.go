package api

import (
	"context"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/google/uuid"
)

type Executor interface {
	Register(ctx context.Context, req models.RegisterRequest) error
	Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error)
	Authorize(ctx context.Context, token string) (*models.User, error)
	TasksList(ctx context.Context, user models.User, filterParams map[string][]string) ([]*models.Task, error)
	TaskByID(ctx context.Context, user models.User, id uuid.UUID) (*models.Task, error)
	CreateTask(ctx context.Context, task models.CreateTaskRequest, user models.User) (uuid.UUID, error)
}
