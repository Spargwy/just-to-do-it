package api

import (
	"net/http"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/Spargwy/just-to-do-it/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func (s *Server) Home(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, world")
}

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

func (s *Server) CreateTask(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("current-client").(models.User)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	req := models.Task{}
	c.Bind(&req)

	//TO-DO implement playground validator
	if req.Title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "field title must be provided")
	}

	err := s.executor.CreateTask(ctx, req, user)
	if err != nil {
		logger.Info("TasksList: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	logger.Info("user %v created task %v", user.ID, req.ID)
	return c.JSON(http.StatusCreated, wrapResponse("task created"))
}
