package models

import (
	"time"

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

type CreateTaskRequest struct {
	ParentTaskID      uuid.NullUUID `json:"parent_task_id" swaggertype:"string"`
	ResponsibleUserID uuid.NullUUID `json:"responsible_user_id" swaggertype:"string"`
	Title             string        `json:"title"`
	Description       *string       `json:"description"`
	TaskGroupID       uuid.NullUUID `json:"task_group_id" swaggertype:"string"`
	Priority          *TaskPriority `json:"priority"`
	EstimateTime      *int          `json:"estimate_time"`
}

type Task struct {
	ID                uuid.UUID     `query:"id" param:"id" json:"id"`
	CreatedAt         time.Time     `query:"created_at" param:"created_at" json:"created_at"`
	ParentTaskID      uuid.NullUUID `query:"parent_task_id" param:"parent_task_id" json:"parentTaskID"`
	CreaterID         uuid.UUID     `query:"creater_id" param:"creater_id" json:"creater_id"`
	ResponsibleUserID uuid.NullUUID `query:"responsible_user_id" param:"responsible_user_id" json:"responsible_user_id"`
	Title             string        `query:"title" param:"title" json:"title"`
	Description       *string       `query:"description" param:"description" json:"description"`
	Status            *TaskStatus   `query:"status" param:"status" json:"status"`
	WorkspaceID       uuid.NullUUID `query:"workspace_id" param:"workspace_id" json:"workspace_id"`
	TaskGroupID       uuid.NullUUID `query:"task_group_id" param:"task_group_id" json:"task_group_id"`
	Priority          *TaskPriority `query:"priority" param:"priority" json:"priority"`
	EstimateTime      *int          `query:"estimate_time" param:"estimate_time" json:"estimate_time"`
	TimeSpent         *int          `query:"time_spent" param:"time_spent" json:"time_spent"`
	DeletedAt         *time.Time    `query:"deleted_at" param:"deleted_at" json:"deleted_at"`
	Archived          bool          `query:"archived" param:"archived" json:"archived"`
}

func (t *Task) Convert(req CreateTaskRequest) {
	t.Title = req.Title
	t.CreatedAt = time.Now()
	if req.Description != nil {
		t.Description = req.Description
	}
	if req.TaskGroupID.Valid {
		t.TaskGroupID = uuid.NullUUID{Valid: true, UUID: req.TaskGroupID.UUID}
	}
	if req.Priority != nil {
		t.Priority = req.Priority
	}
	if req.EstimateTime != nil {
		t.EstimateTime = req.EstimateTime
	}
	if req.ResponsibleUserID.Valid {
		t.ResponsibleUserID = uuid.NullUUID{Valid: true, UUID: req.ResponsibleUserID.UUID}
	}
	if req.ParentTaskID.Valid {
		t.ParentTaskID = uuid.NullUUID{Valid: true, UUID: req.ParentTaskID.UUID}
	}
}

type TaskGroup struct {
	ID    uuid.UUID
	Title string
}
