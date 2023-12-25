package postgres

import (
	"context"
	"fmt"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/Spargwy/just-to-do-it/pkg/config"
	"github.com/Spargwy/just-to-do-it/pkg/db"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ClientPGDB struct {
	db *sqlx.DB
}

func NewPostgres(cfg config.Database) (*ClientPGDB, error) {
	clientDB := ClientPGDB{}

	var err error

	clientDB.db, err = db.NewPostgres(db.DBConfig{
		DBURL: cfg.Client,
	})

	return &clientDB, err
}

func (c *ClientPGDB) Ping(ctx context.Context) error {
	return c.db.Ping()
}

func (c *ClientPGDB) TasksList(whereCondition string, task models.Task) ([]*models.Task, error) {
	var sqlTask Task
	sqlTask.ConvertToSqlStruct(task)
	tasks := []*models.Task{}
	rows := &sqlx.Rows{}
	var err error
	if whereCondition != "" {
		rows, err = c.db.NamedQuery(fmt.Sprintf("select * from tasks where %s", whereCondition), &sqlTask)
	} else {
		rows, err = c.db.NamedQuery("select * from tasks", &sqlTask)
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := models.Task{}
		sqlTask := Task{}
		err := rows.StructScan(&sqlTask)
		if err != nil {
			return nil, fmt.Errorf("structScan: %v", err)
		}
		sqlTask.ConvertFromSqlStruct(&task)
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (c *ClientPGDB) TaskByID(id uuid.UUID) (*models.Task, error) {
	row := Task{}
	var task models.Task
	err := c.db.Get(&row, "select * from tasks where id = $1", id)
	if err != nil {
		return nil, err
	}

	err = row.ConvertFromSqlStruct(&task)
	return &task, err
}

func (c *ClientPGDB) CreateTask(t *models.Task) error {
	task := Task{}
	task.ConvertToSqlStruct(*t)

	tx, err := c.db.Beginx()
	if err != nil {
		return fmt.Errorf("begin tx: %v", err)
	}
	defer tx.Rollback()
	_, err = tx.NamedExec(`
	insert into tasks(
		creater_id,
		created_at,
		parent_task_id,
		responsible_user_id,
		title,
		description,
		status,
		task_group_id,
		priority,
		estimate_time,
		time_spent,
		deleted_at,
		archived
	) values(
		:creater_id,
		:created_at,
		:parent_task_id,
		:responsible_user_id,
		:title,
		:description,
		:status,
		:task_group_id,
		:priority,
		:estimate_time,
		:time_spent,
		:deleted_at,
		:archived
	)`, task)
	if err != nil {
		return err
	}
	err = tx.Commit()

	return err
}

func (c *ClientPGDB) CreateUser(u *models.User) error {
	user := User{}
	user.ConvertToSqlStruct(*u)
	tx, err := c.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("begin tx: %v", err)
	}

	defer tx.Rollback()
	_, err = tx.NamedExec(`
	insert into users(email, name, hashed_password) values(
		:email,
		:name,
		:hashed_password	
	)`, &user)
	if err != nil {
		execError := fmt.Errorf("exec: %v", err)
		err = tx.Rollback()
		if err != nil {
			return fmt.Errorf("rollback: %v", err)
		}

		return execError
	}
	err = tx.Commit()

	return err
}

func (c *ClientPGDB) UserExistsByEmail(email string) (bool, error) {
	var exists bool
	err := c.db.Get(&exists, "select exists(select id from users where email = $1)", email)
	return exists, err
}

func (c *ClientPGDB) GetUserByEmail(email string) (models.User, error) {
	user := User{}
	u := models.User{}
	err := c.db.Get(&user, "select * from users where email = $1", email)
	if err != nil {
		return u, err
	}

	user.ConvertFromSqlStruct(&u)
	return u, err
}

func (c *ClientPGDB) GetUserByID(id uuid.UUID) (models.User, error) {
	user := User{}
	u := models.User{}
	err := c.db.Get(&user, "select * from users where id = $1", id.String())
	if err != nil {
		return u, err
	}
	user.ConvertFromSqlStruct(&u)
	return u, err
}
