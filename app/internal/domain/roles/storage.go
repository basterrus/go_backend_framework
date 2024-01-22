package roles

import "context"

type Storage interface {
	CreateRole(ctx context.Context, role Role) (string, error)
	GetAllRoles(ctx context.Context) ([]Role, error)
	GetRoleByID(ctx context.Context, id string) (Role, error)
	Update(ctx context.Context, role Role) error
	Delete(ctx context.Context, id string) error
}
