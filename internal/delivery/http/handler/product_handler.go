package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/entities"
	"github.com/faisallbhr/light-pos-be/internal/service"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
	"github.com/faisallbhr/light-pos-be/pkg/utils"
	"github.com/faisallbhr/light-pos-be/pkg/validatorx"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
	timeout        time.Duration
}

func NewProductHandler(productService service.ProductService, timeout time.Duration) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		timeout:        timeout,
	}
}

func (h *ProductHandler) CreateOpeningStock(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	var req dto.CreateOpeningStockRequest
	if err := c.ShouldBind(&req); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &req)
		httpx.ResponseError(c, "invalid request body", statusCode, errors)
		return
	}

	imageFile, _ := c.FormFile("image")

	product, err := h.productService.CreateOpeningStock(ctx, &req, imageFile)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, product, "opening stock created successfully", http.StatusOK, nil)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	var params httpx.QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &params)
		httpx.ResponseError(c, "invalid query parameters", statusCode, errors)
		return
	}

	validFields := utils.GetStructFieldNames[entities.Product]()
	if !params.IsValidOrderField(validFields) {
		httpx.ResponseError(c, "invalid order field", http.StatusBadRequest, nil)
		return
	}
	products, total, err := h.productService.GetProducts(ctx, &params)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	meta := httpx.BuildMeta(&params, total)
	httpx.ResponseSuccess(c, products, "products fetched successfully", http.StatusOK, meta)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	id, err := httpx.ParseIDFromParam(c, "id")
	if err != nil {
		httpx.ResponseError(c, "invalid product ID", http.StatusBadRequest, nil)
		return
	}

	product, err := h.productService.GetProductByID(ctx, id)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, product, "product fetched successfully", http.StatusOK, nil)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	id, err := httpx.ParseIDFromParam(c, "id")
	if err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &id)
		httpx.ResponseError(c, "invalid user id", statusCode, errors)
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBind(&req); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &req)
		httpx.ResponseError(c, "invalid request body", statusCode, errors)
		return
	}

	product, err := h.productService.UpdateProduct(ctx, id, &req)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, product, "product updated successfully", http.StatusOK, nil)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	id, err := httpx.ParseIDFromParam(c, "id")
	if err != nil {
		httpx.ResponseError(c, "invalid product ID", http.StatusBadRequest, nil)
		return
	}

	if err := h.productService.DeleteProduct(ctx, id); err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, nil, "product deleted successfully", http.StatusOK, nil)
}
