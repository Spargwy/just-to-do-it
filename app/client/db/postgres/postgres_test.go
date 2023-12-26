package postgres

import (
	"reflect"
	"testing"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/Spargwy/just-to-do-it/pkg/config"
	"github.com/Spargwy/just-to-do-it/pkg/db"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestClientPGDB_TasksList(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		whereCondition string
		task           models.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClientPGDB{
				db: tt.fields.db,
			}
			got, err := c.TasksList(tt.args.whereCondition, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientPGDB.TasksList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientPGDB.TasksList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientPGDB_TaskByID(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClientPGDB{
				db: tt.fields.db,
			}
			got, err := c.TaskByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientPGDB.TaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientPGDB.TaskByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientPGDB_Tasks(t *testing.T) {
	cfg := config.LoadConfig("../../../../")
	clientDB := ClientPGDB{}
	var err error
	clientDB.db, err = db.NewPostgres(db.DBConfig{
		DBURL: cfg.Database.TestDB,
	})
	if err != nil {
		t.Fatalf("NewPostgres error: %v", err)
	}
	t.Cleanup(func() {
		_, err = clientDB.db.Exec("delete from tasks")
		if err != nil {
			t.Fatal(err)
		}
		_, err := clientDB.db.Exec("delete from users")
		if err != nil {
			t.Fatal(err)
		}
	})

	user := models.User{
		Email:          "email",
		HashedPassword: "password",
	}
	err = clientDB.CreateUser(&user)
	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedUser := models.User{
		ID:             user.ID,
		Email:          "email",
		HashedPassword: "password",
	}

	require.Equal(t, expectedUser, user)

	lowPriority := models.LowTaskPriority
	task1 := models.Task{
		CreaterID: user.ID,
		Title:     "title1",
		Priority:  &lowPriority,
	}

	highPriority := models.HighTaskPriority
	task2 := models.Task{
		CreaterID: user.ID,
		Title:     "title2",
		Priority:  &highPriority,
	}
	err = clientDB.CreateTask(&task1)
	if err != nil {
		t.Fatalf("CreateTask: %v", err)
	}

	err = clientDB.CreateTask(&task2)
	if err != nil {
		t.Fatalf("CreateTask: %v", err)
	}

	actualTask, err := clientDB.TaskByID(task1.ID)
	if err != nil {
		t.Fatalf("TaskByID: %v", err)
	}

	status := models.CreatedTaskStatus
	expectedTask1 := models.Task{
		ID:        actualTask.ID,
		CreaterID: user.ID,
		Title:     "title1",
		Priority:  &lowPriority,
		Status:    &status,
	}

	require.Equal(t, &expectedTask1, actualTask)

	expectedTask2 := models.Task{
		ID:        task2.ID,
		CreaterID: user.ID,
		Title:     "title2",
		Priority:  &highPriority,
		Status:    &status,
	}

	filterTask := models.Task{
		Archived: false,
	}

	expectedTasks := []*models.Task{
		&expectedTask1,
		&expectedTask2,
	}

	actualTasks, err := clientDB.TasksList("archived = :archived", filterTask)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, expectedTasks, actualTasks)

}

func TestClientPGDB_Users(t *testing.T) {
	cfg := config.LoadConfig("../../../../")
	clientDB := ClientPGDB{}
	var err error
	clientDB.db, err = db.NewPostgres(db.DBConfig{
		DBURL: cfg.Database.TestDB,
	})
	if err != nil {
		t.Fatalf("NewPostgres error: %v", err)
	}
	t.Cleanup(func() {
		_, err = clientDB.db.Exec("delete from tasks")
		if err != nil {
			t.Fatal(err)
		}
		_, err := clientDB.db.Exec("delete from users")
		if err != nil {
			t.Fatal(err)
		}
	})

	user := models.User{
		Email:          "email",
		HashedPassword: "password",
	}
	err = clientDB.CreateUser(&user)
	if err != nil {
		t.Fatalf(err.Error())
	}

	exists, err := clientDB.UserExistsByEmail(user.Email)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, true, exists)

	userByEmail, err := clientDB.GetUserByEmail(user.Email)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, user, userByEmail)

	userByID, err := clientDB.GetUserByID(user.ID)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, user, userByID)
}

func TestClientPGDB_UserExistsByEmail(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClientPGDB{
				db: tt.fields.db,
			}
			got, err := c.UserExistsByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientPGDB.UserExistsByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClientPGDB.UserExistsByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientPGDB_GetUserByEmail(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClientPGDB{
				db: tt.fields.db,
			}
			got, err := c.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientPGDB.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientPGDB.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
