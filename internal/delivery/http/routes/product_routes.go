package routes

import (
	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/handler"
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup, handler *handler.ProductHandler, db *database.DB) {
	products := rg.Group("/products")
	{
		products.POST("/opening-stock", middleware.PermissionMiddleware("manage_products", db), handler.CreateOpeningStock)
	}
}
