package routes

import (
	"net/http"
	"time"

	"github.com/faisallbhr/light-pos-be/config"
	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/handler"
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/middleware"
	"github.com/faisallbhr/light-pos-be/internal/repository"
	"github.com/faisallbhr/light-pos-be/internal/service"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
	"github.com/faisallbhr/light-pos-be/pkg/jwtx"
	"github.com/gin-gonic/gin"
)

type HandlerRegistry struct {
	AuthHandler *handler.AuthHandler
	UserHandler *handler.UserHandler
}

func SetupRouter(db *database.DB) *gin.Engine {
	cfg := config.LoadConfig()
	jwtManager := jwtx.NewJWTManager(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		cfg.JWT.AccessTTL,
		cfg.JWT.RefreshTTL,
	)

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo, jwtManager)
	authHandler := handler.NewAuthHandler(authService, 10*time.Second)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, 10*time.Second)

	h := &HandlerRegistry{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}

	return InitRoutes(h, jwtManager)
}

func InitRoutes(h *HandlerRegistry, jwtManager *jwtx.JWTManager) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	api := router.Group("/api")

	RegisterAuthRoutes(api, h.AuthHandler)

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtManager))

	RegisterUserRoutes(protected, h.UserHandler)

	router.NoRoute(func(c *gin.Context) {
		httpx.ResponseError(c, "route not found", http.StatusNotFound, nil)
	})

	return router
}
