package authorization

import (
	"context"
	"github.com/basterrus/go_backend_framework/pkg/logging"
)

type service struct {
	storage Storage
	logger  logging.Logger
}

func NewAuthService(authStorage Storage, logger logging.Logger) (Service, error) {
	return &service{
		storage: authStorage,
		logger:  logger,
	}, nil
}

type Service interface {
	SignIn(ctx context.Context) (token string, err error)
	SignUp(ctx context.Context) (token string, err error)
	GetUserByEmail(ctx context.Context, email string) (err error)
	GenerateToken(email, password string) (token string, err error)
}

func (s service) SignIn(ctx context.Context) (token string, err error) {

	return token, nil
}

func (s service) SignUp(ctx context.Context) (token string, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetUserByEmail(ctx context.Context, email string) (err error) {
	//user, err := s.storage.GetUserByEmail(ctx, email)
	//
	return nil
}

func (s service) GenerateToken(email, password string) (token string, err error) {
	user, err := s.storage.GetUserByEmail(context.Background(), email)
	if err != nil {
		return "", err
	}

}
