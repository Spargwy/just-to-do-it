package service

import (
	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/google/uuid"
)

type ClientDB interface {
	TasksList() ([]*models.Task, error)
	TaskByID(id uuid.UUID) (*models.Task, error)
	CreateTask() error
}

type ClientTasker struct {
	db ClientDB
}

func New(db ClientDB) *ClientTasker {
	return &ClientTasker{
		db: db,
	}
}
