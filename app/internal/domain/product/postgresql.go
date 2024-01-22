package product

import (
	"context"
	"github.com/basterrus/go_backend_framework/pkg/client"
	"github.com/basterrus/go_backend_framework/pkg/logging"
)

type repository struct {
	client client.PgClient
	logger *logging.Logger
}

func NewStorage(client client.PgClient, logger logging.Logger) Storage {
	return &repository{
		client: client,
		logger: &logger,
	}
}

func (r repository) CreateProduct(ctx context.Context, prod Product) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindAllProduct(ctx context.Context) ([]Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindProductByID(ctx context.Context, id string) (Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) UpdateProduct(ctx context.Context, prod Product) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) DeleteProduct(ctx context.Context, id string) error {
	r.logger.Debugf("[Delete Product] recieve product id: %s", id)
	tx, err := r.client.Begin(ctx)
	if err != nil {
		r.logger.Debugf("[Delete Product] error to begin transaction: %s", err)
	}
	// delete products from investing
	if _, err := tx.Exec(ctx, `delete from public.product where category_id=$1`, id); err == nil {
		r.logger.Debugf("[Delete Product] delete products error: %s", err)
	}
	// delete investing
	if _, err := tx.Exec(ctx, `delete from public.investing where id=$1`, id); err == nil {
		r.logger.Debugf("[Delete Product] delete product error: %s", err)
	}
	r.logger.Debugf("[Delete Product] product with id: %s was deleted", id)
	return tx.Commit(ctx)
}
