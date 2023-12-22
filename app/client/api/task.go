package api

import (
	"log"
	"net/http"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/Spargwy/just-to-do-it/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Summary
// @Tag Tasks
// @Description list of user tasks
// @Param			Authorization	header		string	true	"Authentication header"
// @Param title query string false "task title"
// @Param id query string false "task id"
// @Param created_at query string false "task created_at"
// @Param parent_task_id query string false "task parent_task_id"
// @Param creater_id query string false "task creater_id"
// @Param responsible_user_id query string false "task responsible_user_id"
// @Param title query string false "task title"
// @Param description query string false "task description"
// @Param status query string false "task status"
// @Param task_group_id query string false "task task_group_id"
// @Param priority query string false "task priority"
// @Param estimate_time query string false "task estimate_time"
// @Param time_spent query string false "task time_spent"
// @Param deleted_at query string false "task deleted_at"
// @Param archived query boolean false "task archived"
// @Success 200
// @Failure 401
// @Failure 500
// @Router /tasks [get]
func (s *Server) TasksList(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("current-client").(models.User)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	filterParams := c.QueryParams()

	tasks, err := s.executor.TasksList(ctx, user, filterParams)
	if err != nil {
		logger.Info("TasksList: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, tasks)
}

// @Summary
// @Tag Tasks
// @Description task by id
// @Param			Authorization	header		string	true	"Authentication header"
// @Param id path string true "task id"
// @Success 200
// @Failure 401
// @Failure 500
// @Router /task/{id} [get]
func (s *Server) TaskByID(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("current-client").(models.User)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "task_id is not uuid")
	}

	tasks, err := s.executor.TaskByID(ctx, user, taskID)
	if err != nil {
		logger.Info("TasksList: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, tasks)
}

// @Summary
// @Tag Tasks
// @Description create task
// @Accept json
// @Param			Authorization	header		string	true	"Authentication header"
// @Param input body models.CreateTaskRequest true "create task body"
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /task [post]
func (s *Server) CreateTask(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("current-client").(models.User)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	req := models.CreateTaskRequest{}
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	log.Print(req.Title)
	//TO-DO implement playground validator
	if req.Title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "field title must be provided")
	}

	taskID, err := s.executor.CreateTask(ctx, req, user)
	if err != nil {
		logger.Info("TasksList: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	logger.Info("user %v created task %v", user.ID, taskID)
	return c.JSON(http.StatusCreated, wrapResponse("task created"))
}
