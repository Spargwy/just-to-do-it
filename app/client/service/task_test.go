package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/google/uuid"
)

func TestClientExecutor_TaskByID(t *testing.T) {
	type fields struct {
		db        ClientDB
		encrypter Encrypter
		jwt       Authenticator
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}

	baseUUID := uuid.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Response
		wantErr bool
	}{
		{
			name: "sql ErrNoRows",
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			fields: fields{
				db: &ClientDBMock{
					TaskByIDFunc: func(id uuid.UUID) (*models.Task, error) {
						return nil, sql.ErrNoRows
					},
				},
			},
			want: models.Response{
				Status:  http.StatusOK,
				Message: "task not found",
			},
			wantErr: false,
		},
		{
			name: "internal error",
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			fields: fields{
				db: &ClientDBMock{
					TaskByIDFunc: func(id uuid.UUID) (*models.Task, error) {
						return nil, errors.New("Internal error")
					},
				},
			},
			want: models.Response{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error",
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			fields: fields{
				db: &ClientDBMock{
					TaskByIDFunc: func(id uuid.UUID) (*models.Task, error) {
						return &models.Task{
							ID: baseUUID,
						}, nil
					},
				},
			},
			want: models.Response{
				Status: http.StatusOK,
				Object: models.Task{
					ID: baseUUID,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClientExecutor{
				db:        tt.fields.db,
				encrypter: tt.fields.encrypter,
				jwt:       tt.fields.jwt,
			}
			got, err := c.TaskByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientExecutor.TaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientExecutor.TaskByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientExecutor_TasksList(t *testing.T) {
	type fields struct {
		db        ClientDB
		encrypter Encrypter
		jwt       Authenticator
	}

	type args struct {
		ctx          context.Context
		filterParams map[string][]string
		filterStruct models.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Response
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: &ClientDBMock{
					TasksListFunc: func(whereCondition string, task models.Task) ([]*models.Task, error) {
						switch whereCondition {
						case "title = :title and archived = :archived":
							return []*models.Task{}, nil
						case "archived = :archived and title = :title":
							return []*models.Task{}, nil
						default:
							t.Fatalf("%s is not valid where condition", whereCondition)
						}
						return []*models.Task{}, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				filterParams: map[string][]string{
					"title":    {"title1"},
					"archived": {"false"},
				},
				filterStruct: models.Task{
					Title:    "title1",
					Archived: false,
				},
			},
			want: models.Response{
				Status: http.StatusOK,
				Object: []*models.Task{},
			},
		},
		{
			name: "internal error response",
			fields: fields{
				db: &ClientDBMock{
					TasksListFunc: func(whereCondition string, task models.Task) ([]*models.Task, error) {
						return nil, errors.New("internal error")
					},
				},
			},
			args: args{
				ctx:          context.Background(),
				filterParams: map[string][]string{},
				filterStruct: models.Task{},
			},
			want: models.Response{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClientExecutor{
				db:        tt.fields.db,
				encrypter: tt.fields.encrypter,
				jwt:       tt.fields.jwt,
			}
			got, err := c.TasksList(tt.args.ctx, tt.args.filterParams, tt.args.filterStruct)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientExecutor.TasksList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientExecutor.TasksList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientExecutor_CreateTask(t *testing.T) {
	baseUUID := uuid.New()
	type fields struct {
		db        ClientDB
		encrypter Encrypter
		jwt       Authenticator
	}
	type args struct {
		ctx  context.Context
		req  models.CreateTaskRequest
		user models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Response
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: &ClientDBMock{
					CreateTaskFunc: func(task *models.Task) error {
						task.ID = baseUUID
						task.CreaterID = baseUUID
						return nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				req: models.CreateTaskRequest{},
				user: models.User{
					ID: baseUUID,
				},
			},
			want: models.Response{
				Status:  http.StatusOK,
				Message: "task created",
				Object:  baseUUID,
			},
		},
		{
			name: "success",
			fields: fields{
				db: &ClientDBMock{
					CreateTaskFunc: func(task *models.Task) error {
						return errors.New("Internal error")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				req: models.CreateTaskRequest{},
				user: models.User{
					ID: baseUUID,
				},
			},
			want: models.Response{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClientExecutor{
				db:        tt.fields.db,
				encrypter: tt.fields.encrypter,
				jwt:       tt.fields.jwt,
			}
			got, err := c.CreateTask(tt.args.ctx, tt.args.req, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientExecutor.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientExecutor.CreateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
