package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/Spargwy/just-to-do-it/pkg/auth/model"
	"github.com/Spargwy/just-to-do-it/pkg/logger"
	"github.com/labstack/echo"
)

var (
	ErrUserNotFound = Errorf("user with this email or password not found")
	ErrInvalidToken = Errorf("incorrect token")
)

type Error struct {
	s string
}

func Errorf(s string, args ...interface{}) Error {
	return Error{s: fmt.Sprintf(s, args...)}
}

func (s *ClientExecutor) Authorize(ctx context.Context, token string) (*models.User, error) {
	claims, err := s.jwt.Parse(token)
	if err != nil {
		logger.Info(err.Error())
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	user, err := s.db.GetUserByID(claims.UserID)
	if err != nil && err == sql.ErrNoRows {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "user not found")
	}

	return &user, err
}

func (s *ClientExecutor) Register(ctx context.Context, req models.RegisterRequest) error {
	exists, err := s.db.UserExistsByEmail(req.Email)
	if err != nil {
		logger.Info("userExistsByEmail: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	if exists {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("user with email %s already registered", req.Email).Error())
	}

	hashed, err := s.encrypter.GenerateHash(req.Password)
	if err != nil {
		logger.Info("generateFromPassword: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	err = s.db.CreateUser(&models.User{
		Email:          req.Email,
		Name:           req.Name,
		HashedPassword: hashed,
	})

	return err
}

func (s *ClientExecutor) Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error) {
	user, err := s.db.GetUserByEmail(req.Email)
	if err != nil && err == sql.ErrNoRows {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "user not found")
	}
	if err != nil {
		logger.Info("GetUserByEmail: %v", err)
		return nil, echo.ErrInternalServerError
	}

	if !s.encrypter.CompareHashAndPassword(req.Password, user.HashedPassword) {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "wrong password")
	}

	token, err := s.jwt.Generate(&model.CustomClaims{
		UserID: user.ID,
	})
	if err != nil {
		logger.Info("generate: %v", err)
		return nil, echo.ErrInternalServerError
	}

	return &models.LoginResponse{Token: token}, nil
}
