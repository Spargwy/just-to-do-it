package service

import (
	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/Spargwy/just-to-do-it/pkg/auth/model"
	"github.com/google/uuid"
)

type ClientDB interface {
	UserExistsByEmail(email string) (bool, error)
	CreateUser(*models.User) error
	GetUserByID(id uuid.UUID) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	TasksList(whereCondition string, user_id uuid.UUID) ([]*models.Task, error)
	TaskByID(id, user_id uuid.UUID) (*models.Task, error)
	CreateTask(task *models.Task) error
}

type Encrypter interface {
	GenerateHash(src string) (string, error)
	CompareHashAndPassword(raw, hashed string) bool
}

type Authenticator interface {
	Generate(claims *model.CustomClaims) (string, error)
	Parse(t string) (*model.CustomClaims, error)
}

type ClientExecutor struct {
	db        ClientDB
	encrypter Encrypter
	jwt       Authenticator
}

func New(db ClientDB, encrypter Encrypter, jwt Authenticator) ClientExecutor {
	return ClientExecutor{
		db:        db,
		encrypter: encrypter,
		jwt:       jwt,
	}
}
