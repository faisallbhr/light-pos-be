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

type AuthHandler struct {
	authService service.AuthService
	timeout     time.Duration
}

func NewAuthHandler(authService service.AuthService, timeout time.Duration) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		timeout:     timeout,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &req)
		httpx.ResponseError(c, "invalid request body", statusCode, errors)
		return
	}

	if err := h.authService.Register(ctx, &req); err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, nil, "user registered successfully", http.StatusOK, nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &req)
		httpx.ResponseError(c, "invalid request body", statusCode, errors)
		return
	}

	user, err := h.authService.Login(ctx, &req)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, user, "user logged in successfully", http.StatusOK, nil)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	var req dto.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &req)
		httpx.ResponseError(c, "invalid request body", statusCode, errors)
		return
	}

	user, err := h.authService.Refresh(ctx, &req)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, user, "token refreshed successfully", http.StatusOK, nil)
}
