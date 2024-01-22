package product

import "context"

type Storage interface {
	CreateProduct(ctx context.Context, prod Product) (string, error)
	FindAllProduct(ctx context.Context) ([]Product, error)
	FindProductByID(ctx context.Context, id string) (Product, error)
	UpdateProduct(ctx context.Context, prod Product) error
	DeleteProduct(ctx context.Context, id string) error
}
