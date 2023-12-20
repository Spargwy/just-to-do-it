package postgres

import (
	"context"
	"fmt"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/Spargwy/just-to-do-it/pkg/config"
	"github.com/Spargwy/just-to-do-it/pkg/db"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type ClientPGDB struct {
	db *pg.DB
}

func NewPostgres(cfg config.Database) (*ClientPGDB, error) {
	clientDB := ClientPGDB{}
	var err error
	clientDB.db, err = db.NewPostgres(db.DBConfig{
		DBURL:           cfg.Client,
		EnableDBLog:     cfg.EnableDBLog,
		EnableLongDBLog: cfg.EnableLongDBLog,
		MaxRetries:      cfg.MaxRetries,
	})

	if err != nil {
		return nil, fmt.Errorf("initialize db: %v", err)
	}

	return &clientDB, err
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
