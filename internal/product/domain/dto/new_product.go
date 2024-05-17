package dto

import (
	"github.com/devkemc/challenge-eulabs/internal/product/domain/entity"
	"github.com/devkemc/challenge-eulabs/pkg/util"
)

type NewProductReq struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int32   `json:"stock"`
	Active      bool    `json:"active"`
}

func (p *NewProductReq) ToEntity() *entity.Product {
	var productEntity entity.Product
	util.Copy(&productEntity, p)
	return &productEntity
}
