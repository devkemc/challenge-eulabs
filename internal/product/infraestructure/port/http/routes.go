package http

import (
	"github.com/devkemc/challenge-eulabs/internal/product/domain/service"
	"github.com/devkemc/challenge-eulabs/internal/product/infraestructure/sqlc"
	"github.com/devkemc/challenge-eulabs/pkg/db/product"
	"github.com/labstack/echo/v4"
)

func Routes(group *echo.Group, query *product.Queries) {
	productRepository := sqlc.NewProductRepositorySqlc(query)
	serviceProduct := service.NewProductService(productRepository)
	handler := NewProductHandler(serviceProduct)
	productRoutes := group.Group("/products")
	{
		productRoutes.GET("", handler.GetProducts)
		productRoutes.POST("", handler.CreateProduct)
		productRoutes.GET("/:id", handler.FindProductById)
		productRoutes.PUT("/:id", handler.UpdateProduct)
		productRoutes.DELETE("/:id", handler.InactivateProduct)
	}
}
