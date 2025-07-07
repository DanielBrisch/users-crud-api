package server

import (
	"users-crud/internal/server/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	routes.RegisterUserRoutes(r, db)
}
