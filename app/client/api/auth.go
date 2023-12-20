package api

import (
	"github.com/labstack/echo"
)

func (s *Server) Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
