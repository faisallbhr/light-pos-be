package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/service"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
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
	if err := c.ShouldBindJSON(&req); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &req)
		httpx.ResponseError(c, "invalid request body", statusCode, errors)
		return
	}

	if err := h.productService.CreateOpeningStock(ctx, &req); err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, nil, "opening stock created successfully", http.StatusOK, nil)
}
