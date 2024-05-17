package sqlc

import (
	"context"
	"github.com/devkemc/challenge-eulabs/internal/product/domain/entity"
	"github.com/devkemc/challenge-eulabs/pkg/db/product"
	"github.com/devkemc/challenge-eulabs/pkg/util"
	"strconv"
)

type ProductRepositorySqlc struct {
	querySqlc *product.Queries
}

func NewProductRepositorySqlc(querySqlc *product.Queries) *ProductRepositorySqlc {
	return &ProductRepositorySqlc{
		querySqlc: querySqlc,
	}
}
func (r *ProductRepositorySqlc) ListProducts(c context.Context) ([]*entity.Product, error) {
	products, err := r.querySqlc.ListProducts(c)
	if err != nil {
		return nil, err
	}
	var productsEntity []*entity.Product
	for _, productModel := range products {
		price, err := strconv.ParseFloat(productModel.PrdPrice, 64)
		if err != nil {
			return nil, err
		}
		productsEntity = append(productsEntity, &entity.Product{
			ID:          productModel.PrdID,
			Name:        productModel.PrdName,
			Description: &productModel.PrdDescription.String,
			Price:       price,
			Stock:       productModel.PrdStock,
		})
	}
	return productsEntity, nil
}
func (r *ProductRepositorySqlc) FindProductById(c context.Context, id uint64) (*entity.Product, error) {
	productModel, err := r.querySqlc.GetProductById(c, id)
	if err != nil {
		return nil, err
	}
	price, err := strconv.ParseFloat(productModel.PrdPrice, 64)
	productEntity := &entity.Product{
		ID:          productModel.PrdID,
		Name:        productModel.PrdName,
		Description: &productModel.PrdDescription.String,
		Price:       price,
		Stock:       productModel.PrdStock,
		Active:      productModel.PrdActive,
	}
	return productEntity, nil
}
func (r *ProductRepositorySqlc) UpdateProduct(c context.Context, productEntity *entity.Product) error {
	return r.querySqlc.UpdateProduct(c, product.UpdateProductParams{
		PrdID:          productEntity.ID,
		PrdName:        productEntity.Name,
		PrdDescription: util.StringToNullString(*productEntity.Description),
		PrdStock:       productEntity.Stock,
		PrdActive:      productEntity.Active,
		PrdPrice:       strconv.FormatFloat(productEntity.Price, 'f', -1, 64),
	})
}
func (r *ProductRepositorySqlc) CreateProduct(c context.Context, productEntity *entity.Product) error {
	return r.querySqlc.InsertProduct(c, product.InsertProductParams{
		PrdName:        productEntity.Name,
		PrdDescription: util.StringToNullString(*productEntity.Description),
		PrdStock:       productEntity.Stock,
		PrdActive:      productEntity.Active,
		PrdPrice:       strconv.FormatFloat(productEntity.Price, 'f', -1, 64),
	})
}
func (r *ProductRepositorySqlc) InactivateProduct(c context.Context, id uint64) error {
	return r.querySqlc.InactivateProduct(c, id)
}
