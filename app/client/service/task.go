package service

import (
	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/google/uuid"
)

func (s *ClientTasker) TasksList() ([]*models.Task, error) {
	return nil, nil
}

func (s *ClientTasker) TaskByID(id uuid.UUID) (*models.Task, error) {
	return nil, nil
}

func (s *ClientTasker) CreateTask() error {
	return nil
}
