package dto

import (
	"github.com/devkemc/challenge-eulabs/internal/product/domain/entity"
	"github.com/devkemc/challenge-eulabs/pkg/util"
)

type ProductAll struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description *string `json:"description"`
	Stock       int32   `json:"stock"`
}
type ListProductsRes struct {
	Products []*ProductAll `json:"products"`
}

func ListResFromEntity(productEntity []*entity.Product) *ListProductsRes {
	var all []*ProductAll
	for _, p := range productEntity {
		var pAll ProductAll
		util.Copy(&pAll, p)
		all = append(all, &pAll)
	}
	return &ListProductsRes{Products: all}
}
