package service

import (
	"context"
	"github.com/devkemc/challenge-eulabs/internal/product/domain/dto"
	"github.com/devkemc/challenge-eulabs/internal/product/domain/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *ProductService {
	return &ProductService{repository}
}

func (s *ProductService) RegisterNewProduct(c context.Context, input *dto.NewProductReq) error {
	productEntity := input.ToEntity()
	return s.repository.CreateProduct(c, productEntity)

}

func (s *ProductService) ListAllProducts(c context.Context) (*dto.ListProductsRes, error) {
	allProducts, err := s.repository.ListProducts(c)
	if err != nil {
		return nil, err
	}
	return dto.ListResFromEntity(allProducts), nil
}

func (s *ProductService) FindProductById(c context.Context, id uint64) (*dto.FindByIdProductRes, error) {
	productEntity, err := s.repository.FindProductById(c, id)
	if err != nil {
		return nil, err
	}
	return dto.FindByIdProductResFromEntity(productEntity), nil
}

func (s *ProductService) UpdateProduct(c context.Context, input *dto.UpdateProductReq, id uint64) (*dto.UpdateProductRes, error) {
	productEntity := input.ToEntity(id)
	if err := s.repository.UpdateProduct(c, productEntity); err != nil {
		return nil, err
	}
	return dto.UpdateProductResFromEntity(productEntity), nil
}

func (s *ProductService) InactivateProducet(c context.Context, id uint64) error {
	return s.repository.InactivateProduct(c, id)

}
