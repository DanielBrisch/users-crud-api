package routes

import (
	"users-crud/internal/handlers"
	"users-crud/internal/middleware"
	"users-crud/internal/repositories"
	usecases "users-crud/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUsecase)

	api := r.Group("/api")
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)

	admin := api.Group("/admin")
	admin.Use(middleware.JWTAuthMiddleware(), middleware.AdminOnly())
	admin.PUT("/users/:id/role", userHandler.UpdateRole)
	admin.DELETE("/:id", userHandler.Delete)

	auth := api.Group("/users")
	auth.Use(middleware.JWTAuthMiddleware())
	auth.GET("/get-all", userHandler.GetAll)
	auth.GET("/:id", userHandler.GetByID)
	auth.PUT("/:id", userHandler.Update)
}
