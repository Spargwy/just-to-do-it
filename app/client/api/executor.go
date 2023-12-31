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
	TasksList(ctx context.Context, filterParams map[string][]string, filterStruct models.Task) (models.Response, error)
	TaskByID(ctx context.Context, id uuid.UUID) (models.Response, error)
	CreateTask(ctx context.Context, task models.CreateTaskRequest, user models.User) (models.Response, error)
}
