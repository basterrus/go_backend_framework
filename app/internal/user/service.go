package user

import "context"

type Service interface {
	Create(ctx context.Context, dto CreateUserDTO) (string, error)
	GetByEmailAndPassword(ctx context.Context, email, password string) (User, error)
	GetOne(ctx context.Context, uuid string) (User, error)
	Update(ctx context.Context, dto UpdateUserDTO) error
	Delete(ctx context.Context, uuid string) error
}
