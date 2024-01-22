package product

import (
	"context"
	"github.com/basterrus/go_backend_framework/pkg/logging"
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
	Create(ctx context.Context, dto CreateCategoryDTO) (string, error)
	GetByID(ctx context.Context, uuid string) (Category, error)
	Update(ctx context.Context, dto UpdateCategoryDTO) error
	Delete(ctx context.Context, id string) error
}

func (s service) Create(ctx context.Context, dto CreateCategoryDTO) (userUUID string, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetByID(ctx context.Context, uuid string) (Category, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(ctx context.Context, dto UpdateCategoryDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.storage.Delete(ctx, id)
}
