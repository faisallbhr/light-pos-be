package httpx

import (
	"net/http"

	"github.com/faisallbhr/light-pos-be/pkg/errorsx"
	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    any               `json:"data,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
	Meta    *Meta             `json:"meta,omitempty"`
}

func ResponseSuccess(c *gin.Context, data any, message string, status int, meta *Meta) {
	c.JSON(status, BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func ResponseError(c *gin.Context, message string, status int, errors map[string]string) {
	c.JSON(status, BaseResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

func HandleServiceError(c *gin.Context, err error) {
	switch errorsx.GetCode(err) {
	case errorsx.ErrNotFound:
		ResponseError(c, err.Error(), http.StatusNotFound, nil)
	case errorsx.ErrUnauthorized:
		ResponseError(c, err.Error(), http.StatusUnauthorized, nil)
	case errorsx.ErrForbidden:
		ResponseError(c, err.Error(), http.StatusForbidden, nil)
	case errorsx.ErrBadRequest:
		ResponseError(c, err.Error(), http.StatusBadRequest, nil)
	case errorsx.ErrConflict:
		ResponseError(c, err.Error(), http.StatusConflict, nil)
	default:
		ResponseError(c, "something went wrong", http.StatusInternalServerError, nil)
	}
}
