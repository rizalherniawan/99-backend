package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizalherniawan/99-backend-test/user-service/config"
	"github.com/rizalherniawan/99-backend-test/user-service/database"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/middleware"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/user"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/user/repository"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/user/service"
)

func main() {
	// for local development
	config.LoadEnv()

	db := database.Init()
	database.RunMigration(db)
	r := gin.Default()

	// middleware for validate API key
	r.Use(middleware.APIKeyAuthMiddleware())

	// middleware for handling error
	r.Use(middleware.ErrorHandler())

	api := r.Group("/users")

	userRepository := repository.New(db)
	userService := service.New(userRepository)

	userHandler := user.New(userService)

	user.RegisterRoutes(api, userHandler)

	r.Run(":8080")
}
