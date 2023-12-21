package api

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/Spargwy/just-to-do-it/pkg/auth/model"
	"github.com/Spargwy/just-to-do-it/pkg/logger"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Authenticator interface {
	Generate(claims *model.CustomClaims) (string, error)
	Parse(t string) (*model.CustomClaims, error)
}

type Server struct {
	router   *echo.Echo
	executor Executor
}

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func wrapResponse(data interface{}) response {
	return response{
		Success: true,
		Data:    data,
	}
}

func New(executor Executor, jwt Authenticator) *Server {
	s := &Server{
		router:   echo.New(),
		executor: executor,
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	s.router.HTTPErrorHandler = errorHandler

	auth := s.router.Group("/auth")
	auth.POST("/register", s.Register)
	auth.POST("/login", s.Login)

	home := s.router.Group("/", s.Authorize)
	home.GET("", func(ctx echo.Context) error { return ctx.JSON(http.StatusOK, "Hello") })

}

func errorHandler(err error, c echo.Context) {
	logrus.Debug(errors.WithStack(err))
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	code := he.Code
	message := he.Message

	if err := c.JSON(code, message); err != nil {
		logger.Debug(errors.WithStack(err))
	}
}

func (s *Server) Start(port string) error {
	if len(port) < 1 {
		logrus.Fatal("Application port is empty")
	}

	addr := net.JoinHostPort("", port)

	// create a new server
	server := &http.Server{
		Addr:         addr,              // configure the bind address
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	return s.router.StartServer(server)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}
