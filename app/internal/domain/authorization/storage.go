package authorization

import (
	"context"
	"github.com/basterrus/go_backend_framework/internal/domain/user"
)

type Storage interface {
	SignIn(ctx context.Context) (token string, err error)
	SignUp()
	GetUserByEmail(ctx context.Context, email string) (user user.User, err error)
}
