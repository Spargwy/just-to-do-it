package service

import (
	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/google/uuid"
)

func (s *ClientExecutor) TasksList() ([]*models.Task, error) {
	return nil, nil
}

func (s *ClientExecutor) TaskByID(id uuid.UUID) (*models.Task, error) {
	return nil, nil
}

func (s *ClientExecutor) CreateTask() error {
	return nil
}
