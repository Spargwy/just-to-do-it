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
	ID                uuid.UUID      `param:"id" json:"id"`
	CreatedAt         time.Time      `param:"created_at" json:"created_at"`
	ParentTaskID      uuid.NullUUID  `param:"parent_task_id" json:"parentTaskID"`
	CreaterID         uuid.UUID      `param:"creater_id" json:"creater_id"`
	ResponsibleUserID uuid.NullUUID  `param:"responsible_user_id" json:"responsible_user_id"`
	Title             string         `param:"title" json:"title"`
	Description       *string        `param:"description" json:"description"`
	Status            *string        `param:"status" json:"status"`
	TaskGroupID       *uuid.NullUUID `param:"task_group_id" json:"task_group_id"`
	Priority          *TaskPriority  `param:"priority" json:"priority"`
	EstimateTime      *int           `param:"estimate_time" json:"estimate_time"`
	TimeSpent         *int           `param:"time_spent" json:"time_spent"`
	DeletedAt         *time.Time     `param:"deleted_at" json:"deleted_at"`
	Archived          bool           `param:"archived" json:"archived"`
}

type TaskGroup struct {
	ID    uuid.UUID
	Title string
}
