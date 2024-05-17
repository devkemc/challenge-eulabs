package dto

import (
	"github.com/devkemc/challenge-eulabs/internal/product/domain/entity"
	"github.com/devkemc/challenge-eulabs/pkg/util"
)

type FindByIdProductRes struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description *string `json:"description"`
	Stock       int32   `json:"stock"`
	Active      bool    `json:"active"`
}

func FindByIdProductResFromEntity(productEntity *entity.Product) *FindByIdProductRes {
	var res FindByIdProductRes
	util.Copy(&res, productEntity)
	return &res

}
