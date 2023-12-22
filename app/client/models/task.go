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
	ParentTaskID      *uuid.UUID     `json:"parent_task_id" swaggertype:"string"`
	ResponsibleUserID *uuid.UUID     `json:"responsible_user_id" swaggertype:"string"`
	Title             string         `json:"title"`
	Description       *string        `json:"description"`
	TaskGroupID       *uuid.NullUUID `json:"task_group_id" swaggertype:"string"`
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

func (t *Task) Convert(req CreateTaskRequest) {
	t.Title = req.Title
	t.CreatedAt = time.Now()
	if req.Description != nil {
		t.Description = req.Description
	}
	if req.TaskGroupID != nil {
		t.TaskGroupID = req.TaskGroupID
	}
	if req.Priority != nil {
		t.Priority = req.Priority
	}
	if req.EstimateTime != nil {
		t.EstimateTime = req.EstimateTime
	}
	if req.ResponsibleUserID != nil {
		t.ResponsibleUserID = uuid.NullUUID{Valid: true, UUID: *req.ResponsibleUserID}
	}
	if req.ParentTaskID != nil {
		t.ParentTaskID = uuid.NullUUID{Valid: true, UUID: *req.ParentTaskID}
	}
}

type TaskGroup struct {
	ID    uuid.UUID
	Title string
}
