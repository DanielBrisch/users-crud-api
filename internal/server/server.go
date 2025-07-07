package server

import (
	_ "users-crud/docs"

	"users-crud/internal/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	r := gin.New()

	appLogger := logger.InitLogger()

	r.Use(RequestLogger(appLogger))

	r.Use(MiddlewareLogger())
	r.Use(MiddlewareRecovery())
	r.Use(MiddlewareCORS())
	r.Use(MiddlewareRateLimit())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	RegisterRoutes(r, db)

	return r
}
