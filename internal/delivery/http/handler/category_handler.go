package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/faisallbhr/light-pos-be/internal/entities"
	"github.com/faisallbhr/light-pos-be/internal/service"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
	"github.com/faisallbhr/light-pos-be/pkg/utils"
	"github.com/faisallbhr/light-pos-be/pkg/validatorx"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService service.CategoryService
	timeout         time.Duration
}

func NewCategoryHandler(categoryService service.CategoryService, timeout time.Duration) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
		timeout:         timeout,
	}
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	var params httpx.QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &params)
		httpx.ResponseError(c, "invalid query parameters", statusCode, errors)
		return
	}

	validFields := utils.GetStructFieldNames[entities.Category]()
	if !params.IsValidOrderField(validFields) {
		httpx.ResponseError(c, "invalid order field", http.StatusBadRequest, nil)
		return
	}

	categories, total, err := h.categoryService.GetCategories(ctx, &params)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	meta := httpx.BuildMeta(&params, total)
	httpx.ResponseSuccess(c, categories, "categories retrieved successfully", http.StatusOK, meta)
}
