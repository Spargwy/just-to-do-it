package api

import (
	"net/http"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/labstack/echo"
)

const (
	authorizationHeader = "Authorization"
)

func (s *Server) Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		token := c.Request().Header.Get(authorizationHeader)

		client, err := s.executor.Authorize(ctx, token)
		if err != nil {
			return err
		}

		if client == nil {
			return c.JSON(http.StatusUnauthorized, "wrong email or password")
		}

		c.Set("current-client", *client)

		return next(c)
	}
}

func (s *Server) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := models.RegisterRequest{}

	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = s.executor.Register(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, wrapResponse("created"))
}

func (s *Server) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := models.LoginRequest{}

	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	data, err := s.executor.Login(ctx, req)
	if err != nil {
		return err
	}

	if data == nil {
		return c.JSON(http.StatusUnauthorized, "wrong email or password")
	}

	return c.JSON(http.StatusOK, data)
}
