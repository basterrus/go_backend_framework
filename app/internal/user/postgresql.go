package user

import (
	"context"
	"github.com/basterrus/go_backend_framework/pkg/client"
	"github.com/basterrus/go_backend_framework/pkg/logging"
)

type repository struct {
	client client.PgClient
	logger *logging.Logger
}

func (r *repository) FindAll(ctx context.Context) (u []User, err error) {

	return u, nil
}
