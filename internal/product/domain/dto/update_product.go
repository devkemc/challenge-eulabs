package dto

import (
	"github.com/devkemc/challenge-eulabs/internal/product/domain/entity"
	"github.com/devkemc/challenge-eulabs/pkg/util"
)

type UpdateProductReq struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int32   `json:"stock"`
	Active      bool    `json:"active"`
}

func (p *UpdateProductReq) ToEntity(id uint64) *entity.Product {
	var productEntity entity.Product
	util.Copy(&productEntity, p)
	productEntity.ID = id
	return &productEntity
}

type UpdateProductRes struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int32   `json:"stock"`
	Active      bool    `json:"active"`
}

func UpdateProductResFromEntity(productEntity *entity.Product) *UpdateProductRes {
	var productRes UpdateProductRes
	util.Copy(&productRes, productEntity)
	return &productRes
}
