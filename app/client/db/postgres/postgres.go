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

func (c *ClientPGDB) TasksList(whereCondition string, user_id uuid.UUID) ([]*models.Task, error) {
	row := []*models.Task{}
	if whereCondition != "" {
		err := c.db.Model(&row).
			Where(whereCondition).
			Where("creater_id = ?", user_id).
			Select()
		return row, err
	} else {
		err := c.db.Model(&row).Where("creater_id = ?", user_id).Select()
		return row, err
	}
}

func (c *ClientPGDB) TaskByID(id uuid.UUID, user_id uuid.UUID) (*models.Task, error) {
	row := models.Task{}
	err := c.db.Model(&row).Where("id = ? and creater_id = ?", id, user_id).Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &row, err
}

func (c *ClientPGDB) CreateTask(task *models.Task) error {
	return c.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		_, err := tx.Model(task).Insert()

		return err
	})
}

func (c *ClientPGDB) CreateUser(user *models.User) error {
	return c.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		_, err := c.db.Model(user).Insert()
		return err
	})
}

func (c *ClientPGDB) UserExistsByEmail(email string) (bool, error) {
	var row models.User
	exists, err := c.db.Model(&row).Where("email = ?", email).Exists()
	return exists, err
}

func (c *ClientPGDB) GetUserByEmail(email string) (models.User, error) {
	var row models.User
	err := c.db.Model(&row).Where("email = ?", email).Select()
	return row, err
}

func (c *ClientPGDB) GetUserByID(id uuid.UUID) (models.User, error) {
	var row models.User
	err := c.db.Model(&row).Where("id = ?", id).Select()
	return row, err
}
