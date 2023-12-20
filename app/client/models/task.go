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
	ParentTaskID      *uuid.UUID     `json:"parent_task_id"`
	ResponsibleUserID *uuid.UUID     `json:"responsible_user_id"`
	Title             string         `json:"title"`
	Description       *string        `json:"description"`
	TaskGroupID       *uuid.NullUUID `json:"task_group_id"`
	Priority          *TaskPriority  `json:"priority"`
	EstimateTime      *int           `json:"estimate_time"`
}

type Task struct {
	ID                uuid.UUID
	ParentTaskID      uuid.NullUUID
	CreaterID         uuid.UUID
	ResponsibleUserID uuid.NullUUID
	Title             string
	Description       *string
	Status            *string
	TaskGroupID       *uuid.NullUUID
	Priority          *TaskPriority
	EstimateTime      *int
	TimeSpent         *int
	DeletedAt         *time.Time
	Archived          bool
}

type TaskGroup struct {
	ID    uuid.UUID
	Title string
}
