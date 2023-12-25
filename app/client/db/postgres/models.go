package postgres

import (
	"database/sql"
	"time"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/google/uuid"
)

type TaskStatus string

const (
	CreatedTaskStatus    TaskStatus = "CREATED"
	InProgressTaskStatus TaskStatus = "IN_PROGRESS"
	TestingTaskStatus    TaskStatus = "TESTING"
	DoneTaskStatus       TaskStatus = "DONE"
)

type TaskPriority string

const (
	LowTaskPriority    TaskPriority = "LOW"
	MediumTaskPriority TaskPriority = "MEDIUM"
	HighTaskPriority   TaskPriority = "HIGH"
	ExtraTaskPriority  TaskPriority = "EXTRA"
)

type Task struct {
	ID                sql.NullString `db:"id"`
	CreatedAt         time.Time      `db:"created_at"`
	ParentTaskID      sql.NullString `db:"parent_task_id"`
	CreaterID         string         `db:"creater_id"`
	ResponsibleUserID sql.NullString `db:"responsible_user_id"`
	Title             string         `db:"title"`
	Description       sql.NullString `db:"description"`
	Status            string         `db:"status"`
	TaskGroupID       sql.NullString `db:"task_group_id"`
	WorkspaceID       sql.NullString `db:"workspace_id"`
	Priority          TaskPriority   `db:"priority"`
	EstimateTime      sql.NullInt64  `db:"estimate_time"`
	TimeSpent         sql.NullInt64  `db:"time_spent"`
	DeletedAt         sql.NullTime   `db:"deleted_at"`
	Archived          bool           `db:"archived"`
}

func (t *Task) ConvertToSqlStruct(task models.Task) {
	t.CreatedAt = task.CreatedAt
	t.Title = task.Title
	t.CreaterID = task.CreaterID.String()
	t.Archived = task.Archived

	if task.ParentTaskID.Valid {
		t.ParentTaskID = sql.NullString{Valid: true, String: task.ParentTaskID.UUID.String()}
	}
	if task.ResponsibleUserID.Valid {
		t.ResponsibleUserID = sql.NullString{Valid: true, String: task.ResponsibleUserID.UUID.String()}
	}
	if task.WorkspaceID.Valid {
		t.WorkspaceID = sql.NullString{Valid: true, String: task.WorkspaceID.UUID.String()}
	}
	if task.Description != nil {
		t.Description = sql.NullString{Valid: true, String: *task.Description}
	}
	if task.Status != nil {
		t.Status = string(*task.Status)
	} else {
		t.Status = string(CreatedTaskStatus)
	}
	if task.TaskGroupID.Valid {
		t.TaskGroupID = sql.NullString{Valid: true, String: task.TaskGroupID.UUID.String()}
	}
	if task.Priority != nil {
		t.Priority = TaskPriority(*task.Priority)
	}
	if task.EstimateTime != nil {
		t.EstimateTime = sql.NullInt64{Valid: true, Int64: int64(*task.EstimateTime)}
	}
	if task.TimeSpent != nil {
		t.TimeSpent = sql.NullInt64{Valid: true, Int64: int64(*task.TimeSpent)}
	}
	if task.DeletedAt != nil {
		t.DeletedAt = sql.NullTime{Valid: true, Time: (*task.DeletedAt)}
	}
}

func (t *Task) ConvertFromSqlStruct(task *models.Task) error {
	task.CreaterID = uuid.MustParse(t.CreaterID)
	task.CreatedAt = t.CreatedAt
	task.Title = t.Title
	if t.ID.Valid {
		id, err := uuid.Parse(t.ID.String)
		if err != nil {
			return err
		}
		task.ID = id
	}
	if t.ParentTaskID.Valid {
		id, err := uuid.Parse(t.ParentTaskID.String)
		if err != nil {
			return err
		}
		task.ParentTaskID = uuid.NullUUID{Valid: true, UUID: id}
	}

	if t.ResponsibleUserID.Valid {
		id, err := uuid.Parse(t.ResponsibleUserID.String)
		if err != nil {
			return err
		}
		task.ResponsibleUserID = uuid.NullUUID{Valid: true, UUID: id}
	}
	if t.Description.Valid {
		task.Description = &t.Description.String
	}

	status := models.TaskStatus(t.Status)
	task.Status = &status

	if t.TaskGroupID.Valid {
		id, err := uuid.Parse(t.TaskGroupID.String)
		if err != nil {
			return err
		}
		task.TaskGroupID = uuid.NullUUID{Valid: true, UUID: id}
	}

	priority := models.TaskPriority(t.Priority)
	task.Priority = &priority

	if t.EstimateTime.Valid {
		estimate := int(t.EstimateTime.Int64)
		task.EstimateTime = &estimate
	}
	if t.TimeSpent.Valid {
		spent := int(t.TimeSpent.Int64)
		task.TimeSpent = &spent
	}
	if t.DeletedAt.Valid {
		task.DeletedAt = &t.DeletedAt.Time
	}
	task.Archived = t.Archived

	t.CreatedAt = task.CreatedAt
	t.Title = task.Title
	t.CreaterID = task.CreaterID.String()
	t.Archived = task.Archived

	if task.ParentTaskID.Valid {
		t.ParentTaskID = sql.NullString{Valid: true, String: task.ParentTaskID.UUID.String()}
	}
	if task.ResponsibleUserID.Valid {
		t.ResponsibleUserID = sql.NullString{Valid: true, String: task.ResponsibleUserID.UUID.String()}
	}
	if task.Description != nil {
		t.Description = sql.NullString{Valid: true, String: *task.Description}
	}
	if task.Status != nil {
		t.Status = string(*task.Status)
	} else {
		t.Status = string(CreatedTaskStatus)
	}
	if task.TaskGroupID.Valid {
		t.TaskGroupID = sql.NullString{Valid: true, String: task.TaskGroupID.UUID.String()}
	}
	if task.Priority != nil {
		t.Priority = TaskPriority(*task.Priority)
	}
	if task.EstimateTime != nil {
		t.EstimateTime = sql.NullInt64{Valid: true, Int64: int64(*task.EstimateTime)}
	}
	if task.TimeSpent != nil {
		t.TimeSpent = sql.NullInt64{Valid: true, Int64: int64(*task.TimeSpent)}
	}
	if task.DeletedAt != nil {
		t.DeletedAt = sql.NullTime{Valid: true, Time: (*task.DeletedAt)}
	}

	return nil
}

type User struct {
	ID             string         `json:"id" db:"id"`
	Email          string         `json:"email" db:"email"`
	Name           sql.NullString `json:"name" db:"name"`
	HashedPassword string         `json:"hashed_password" db:"hashed_password"`
}

func (u *User) ConvertToSqlStruct(user models.User) {
	u.ID = user.Email
	u.Email = user.Email
	u.HashedPassword = user.HashedPassword
	if user.Name != nil {
		u.Name = sql.NullString{Valid: true, String: user.Email}
	}
}

func (u *User) ConvertFromSqlStruct(user *models.User) {
	user.ID = uuid.MustParse(u.ID)
	user.Email = u.Email
	user.HashedPassword = u.HashedPassword
	if u.Name.Valid {
		user.Name = &u.Name.String
	}
}
