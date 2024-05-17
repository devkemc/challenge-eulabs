package http

import (
	"github.com/devkemc/challenge-eulabs/internal/product/domain/dto"
	"github.com/devkemc/challenge-eulabs/internal/product/domain/service"
	"github.com/devkemc/challenge-eulabs/pkg/response"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}
func (h *ProductHandler) GetProducts(c echo.Context) error {
	output, err := h.service.ListAllProducts(c.Request().Context())
	if err != nil {
		log.Print(err)
		return response.Error(c, http.StatusInternalServerError, err, "failed to list products")
	}
	return response.JSON(c, http.StatusOK, output)
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var input dto.NewProductReq
	if err := c.Bind(&input); err != nil {
		log.Print(err)
		return response.Error(c, http.StatusBadRequest, err, "failed to bind input")
	}
	if err := h.service.RegisterNewProduct(c.Request().Context(), &input); err != nil {
		log.Print(err)
		return response.Error(c, http.StatusInternalServerError, err, "failed to create product")
	}
	return response.JSON(c, http.StatusCreated, nil)
}

func (h *ProductHandler) GetProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Print(err)
		return response.Error(c, http.StatusBadRequest, err, "failed to parse id")
	}
	output, errService := h.service.FindProductById(c.Request().Context(), id)
	if errService != nil {
		log.Print(errService)
		return response.Error(c, http.StatusInternalServerError, err, "failed to find product")
	}
	return response.JSON(c, http.StatusOK, output)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Print(err)
		return response.Error(c, http.StatusBadRequest, err, "failed to parse id")
	}
	var input dto.UpdateProductReq
	if err := c.Bind(&input); err != nil {
		log.Print(err)
		return response.Error(c, http.StatusBadRequest, err, "failed to bind input")
	}
	output, err := h.service.UpdateProduct(c.Request().Context(), &input, id)
	if err != nil {
		log.Print(err)
		return response.Error(c, http.StatusInternalServerError, err, "failed to update product")
	}
	return response.JSON(c, http.StatusOK, output)
}

func (h *ProductHandler) InactivateProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Print(err)
		return response.Error(c, http.StatusBadRequest, err, "failed to parse id")
	}
	if err := h.service.InactivateProducet(c.Request().Context(), id); err != nil {
		log.Print(err)
		return response.Error(c, http.StatusInternalServerError, err, "failed to inactivate product")
	}
	return response.JSON(c, http.StatusOK, nil)
}
