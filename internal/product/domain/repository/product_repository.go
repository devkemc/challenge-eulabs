package repository

import (
	"context"
	"github.com/devkemc/challenge-eulabs/internal/product/domain/entity"
)

type ProductRepository interface {
	ListProducts(c context.Context) ([]*entity.Product, error)
	FindProductById(c context.Context, id uint64) (*entity.Product, error)
	UpdateProduct(c context.Context, product *entity.Product) error
	CreateProduct(c context.Context, product *entity.Product) error
	InactivateProduct(c context.Context, id uint64) error
}
