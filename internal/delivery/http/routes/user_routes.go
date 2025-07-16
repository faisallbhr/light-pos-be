package routes

import (
	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/handler"
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, handler *handler.UserHandler, db *database.DB) {
	users := rg.Group("/users")
	{
		users.GET("/", middleware.PermissionMiddleware("manage_users", db), handler.GetAllUsers)
		users.GET("/me", handler.Me)
		users.GET("/:id", middleware.PermissionMiddleware("manage_users", db), handler.GetUserByID)
		users.PATCH("/:id", handler.UpdateUser)
		users.PATCH("/:id/password", handler.ChangePassword)
		users.DELETE("/:id", middleware.PermissionMiddleware("manage_users", db), handler.DeleteUser)
	}
}
