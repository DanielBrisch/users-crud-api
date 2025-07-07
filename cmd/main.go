package main

import (
	"log"
	"users-crud/internal/config"
	"users-crud/internal/models"
	"users-crud/internal/server"
)

// @title           Users API
// @version         1.0
// @description     User management API with JWT and roles.
// @termsOfService  http://localhost

// @contact.name   Daniel Dev
// @contact.email  daniel@example.com

// @host      localhost:8080
// @BasePath  /api
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	db.AutoMigrate(&models.User{})

	r := server.Router(db)
	r.Run(":8080")
}
