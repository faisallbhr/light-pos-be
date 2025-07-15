package routes

import (
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, handler *handler.UserHandler) {
	users := rg.Group("/users")
	{
		users.GET("/", handler.GetAllUsers)
		users.GET("/me", handler.Me)
		users.GET("/:id", handler.GetUserByID)
		users.PATCH("/:id", handler.UpdateUser)
		users.PATCH("/:id/password", handler.ChangePassword)
		users.DELETE("/:id", handler.DeleteUser)
	}
}
