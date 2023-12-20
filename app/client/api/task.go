package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) Home(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, world")
}
