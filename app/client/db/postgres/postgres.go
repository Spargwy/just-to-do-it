package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type ClientPGDB struct {
	db *pg.DB
}

func NewPostgres(dbURL string) (*ClientPGDB, error) {
	log.Print(dbURL)
	opts, err := pg.ParseURL(dbURL)
	if err != nil {
		return nil, fmt.Errorf("parseURL: %v", err)
	}

	clientDB := ClientPGDB{
		db: pg.Connect(opts),
	}

	return &clientDB, nil
}

func (c *ClientPGDB) Ping(ctx context.Context) error {
	return c.db.Ping(ctx)
}

func (c *ClientPGDB) ClientTasks() ([]*models.Task, error) {
	return nil, nil
}

func (db *ClientPGDB) TasksList() ([]*models.Task, error) {
	return nil, nil
}

func (db *ClientPGDB) TaskByID(id uuid.UUID) (*models.Task, error) {
	return nil, nil
}

func (db *ClientPGDB) CreateTask() error {
	return nil
}
