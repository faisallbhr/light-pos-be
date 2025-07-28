package routes

import (
	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/handler"
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, handler *handler.CategoryHandler, db *database.DB) {
	categories := rg.Group("/categories")
	{
		categories.GET("/", middleware.PermissionMiddleware("manage_products", db), handler.GetCategories)
	}
}
