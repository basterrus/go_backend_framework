package roles

import (
	"context"
	"errors"
	"fmt"
	"github.com/basterrus/go_backend_framework/internal/apperror"
	"github.com/basterrus/go_backend_framework/pkg/logging"
	"github.com/google/uuid"
)

type service struct {
	storage Storage
	logger  logging.Logger
}

func NewService(userStorage Storage, logger logging.Logger) (Service, error) {
	return &service{
		storage: userStorage,
		logger:  logger,
	}, nil
}

type Service interface {
	Create(ctx context.Context, dto CreateUserDTO) (string, error)
	GetByUUID(ctx context.Context, uuid string) (User, error)
	GetOne(ctx context.Context, uuid string) (User, error)
	Update(ctx context.Context, dto UpdateUserDTO) error
	Delete(ctx context.Context, uuid string) error
}

func (s service) GetOne(ctx context.Context, uuid string) (u User, err error) {
	u, err = s.storage.FindOne(ctx, uuid)

	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return u, err
		}
		return u, fmt.Errorf("failed to find user by uuid. error: %w", err)
	}
	return u, nil
}

func (s service) Create(ctx context.Context, dto CreateUserDTO) (userUUID string, err error) {
	s.logger.Debug("check password and repeat password")
	if dto.Password != dto.RepeatPassword {
		return userUUID, apperror.BadRequestError("password does not match repeat password")
	}
	nu := NewUser(dto)
	nu.Uuid = uuid.New()
	s.logger.Debug("generate password hash")
	err = nu.GeneratePasswordHash()
	if err != nil {
		s.logger.Errorf("failed to create user due to error %v", err)
		return
	}
	return s.storage.Create(ctx, nu)
}

func (s service) GetByUUID(ctx context.Context, uuid string) (user User, err error) {
	user, err = s.storage.FindByUUID(ctx, uuid)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s service) Update(ctx context.Context, dto UpdateUserDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete(ctx context.Context, uuid string) error {
	err := s.storage.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}
