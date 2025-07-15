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

type UserHandler struct {
	userService service.UserService
	timeout     time.Duration
}

func NewUserHandler(userService service.UserService, timeout time.Duration) *UserHandler {
	return &UserHandler{
		userService: userService,
		timeout:     timeout,
	}
}

func (h *UserHandler) Me(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	userID := c.GetUint("user_id")
	if userID == 0 {
		httpx.ResponseError(c, "unauthorized", http.StatusUnauthorized, nil)
		return
	}

	user, err := h.userService.Me(ctx, userID)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, user, "user fetched successfully", http.StatusOK, nil)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	var params httpx.QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &params)
		httpx.ResponseError(c, "invalid query parameters", statusCode, errors)
		return
	}

	users, total, err := h.userService.GetAllUsers(ctx, &params)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	meta := httpx.BuildMeta(&params, total)
	httpx.ResponseSuccess(c, users, "users fetched successfully", http.StatusOK, meta)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	id, err := httpx.ParseIDFromParam(c, "id")
	if err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &id)
		httpx.ResponseError(c, "invalid user id", statusCode, errors)
		return
	}

	user, err := h.userService.GetUserByID(ctx, id)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, user, "user fetched successfully", http.StatusOK, nil)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	id, err := httpx.ParseIDFromParam(c, "id")
	if err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &id)
		httpx.ResponseError(c, "invalid user id", statusCode, errors)
		return
	}

	var req dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &req)
		httpx.ResponseError(c, "invalid request body", statusCode, errors)
		return
	}

	user, err := h.userService.UpdateUser(ctx, id, &req)
	if err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, user, "user updated successfully", http.StatusOK, nil)
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	id, err := httpx.ParseIDFromParam(c, "id")
	if err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &id)
		httpx.ResponseError(c, "invalid user id", statusCode, errors)
		return
	}

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &req)
		httpx.ResponseError(c, "invalid request body", statusCode, errors)
		return
	}

	if err := h.userService.ChangePassword(ctx, id, &req); err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, nil, "password changed successfully", http.StatusOK, nil)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	id, err := httpx.ParseIDFromParam(c, "id")
	if err != nil {
		errors, statusCode := validatorx.TranslateErrorMessage(err, &id)
		httpx.ResponseError(c, "invalid user id", statusCode, errors)
		return
	}

	if err := h.userService.DeleteUser(ctx, id); err != nil {
		httpx.HandleServiceError(c, err)
		return
	}

	httpx.ResponseSuccess(c, nil, "user deleted successfully", http.StatusOK, nil)
}
