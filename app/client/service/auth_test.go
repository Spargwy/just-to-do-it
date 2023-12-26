package service

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/Spargwy/just-to-do-it/app/client/models"
	"github.com/Spargwy/just-to-do-it/pkg/auth/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestClientExecutor_Register(t *testing.T) {
	baseEmail := "email"
	type fields struct {
		db        ClientDB
		encrypter Encrypter
		jwt       Authenticator
	}
	type args struct {
		ctx context.Context
		req models.RegisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: models.RegisterRequest{
					Email:    baseEmail,
					Password: "password",
				},
			},
			fields: fields{
				db: &ClientDBMock{
					UserExistsByEmailFunc: func(email string) (bool, error) {
						require.Equal(t, baseEmail, email)
						return false, nil
					},
					CreateUserFunc: func(user *models.User) error {
						require.Equal(t, baseEmail, user.Email)
						require.Equal(t, "hashed", user.HashedPassword)
						return nil
					},
				},
				encrypter: &EncrypterMock{
					GenerateHashFunc: func(src string) (string, error) {
						return "hashed", nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "exists",
			args: args{
				ctx: context.Background(),
				req: models.RegisterRequest{
					Email:    baseEmail,
					Password: "password",
				},
			},
			fields: fields{
				db: &ClientDBMock{
					UserExistsByEmailFunc: func(email string) (bool, error) {
						require.Equal(t, baseEmail, email)
						return true, nil
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ClientExecutor{
				db:        tt.fields.db,
				encrypter: tt.fields.encrypter,
				jwt:       tt.fields.jwt,
			}
			if err := s.Register(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("ClientExecutor.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClientExecutor_Authorize(t *testing.T) {
	baseUUID := uuid.New()

	type fields struct {
		db        ClientDB
		encrypter Encrypter
		jwt       Authenticator
	}
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: &ClientDBMock{
					GetUserByIDFunc: func(id uuid.UUID) (models.User, error) {
						require.Equal(t, baseUUID, id)
						return models.User{
							ID: baseUUID,
						}, nil
					},
				},
				jwt: &AuthenticatorMock{
					ParseFunc: func(t string) (*model.CustomClaims, error) {
						return &model.CustomClaims{
							UserID: baseUUID,
						}, nil
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				token: "token",
			},
			want: &models.User{
				ID: baseUUID,
			},
		},
		{
			name: "user not found",
			fields: fields{
				db: &ClientDBMock{
					GetUserByIDFunc: func(id uuid.UUID) (models.User, error) {
						require.Equal(t, baseUUID, id)
						return models.User{}, sql.ErrNoRows
					},
				},
				jwt: &AuthenticatorMock{
					ParseFunc: func(t string) (*model.CustomClaims, error) {
						return &model.CustomClaims{
							UserID: baseUUID,
						}, nil
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				token: "token",
			},
			wantErr: true,
		},
		{
			name: "invalid token",
			fields: fields{
				jwt: &AuthenticatorMock{
					ParseFunc: func(t string) (*model.CustomClaims, error) {
						return nil, errors.New("invalid token")
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				token: "token",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ClientExecutor{
				db:        tt.fields.db,
				encrypter: tt.fields.encrypter,
				jwt:       tt.fields.jwt,
			}
			got, err := s.Authorize(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientExecutor.Authorize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientExecutor.Authorize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientExecutor_Login(t *testing.T) {
	type fields struct {
		db        ClientDB
		encrypter Encrypter
		jwt       Authenticator
	}
	type args struct {
		ctx context.Context
		req models.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.LoginResponse
		wantErr bool
	}{
		{
			args: args{
				ctx: context.Background(),
				req: models.LoginRequest{
					Email:    "email",
					Password: "password",
				},
			},
			fields: fields{
				db: &ClientDBMock{
					GetUserByEmailFunc: func(email string) (models.User, error) {
						require.Equal(t, "email", email)
						return models.User{
							Email:          email,
							HashedPassword: "hashed",
						}, nil
					},
				},
				encrypter: &EncrypterMock{
					CompareHashAndPasswordFunc: func(raw, hashed string) bool {
						require.Equal(t, "hashed", hashed)
						return true
					},
				},
				jwt: &AuthenticatorMock{
					GenerateFunc: func(claims *model.CustomClaims) (string, error) {
						return "token", nil
					},
				},
			},
			want: &models.LoginResponse{
				Token: "token",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ClientExecutor{
				db:        tt.fields.db,
				encrypter: tt.fields.encrypter,
				jwt:       tt.fields.jwt,
			}
			got, err := s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientExecutor.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientExecutor.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
