package routes

import (
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *handler.AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/refresh", handler.Refresh)
	}
}
