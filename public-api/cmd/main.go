package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizalherniawan/99-backend-test/public-api/config"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/listing"
	listingService "github.com/rizalherniawan/99-backend-test/public-api/internal/listing/service"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/middleware"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/user"
	userService "github.com/rizalherniawan/99-backend-test/public-api/internal/user/service"
)

func main() {
	// for local development
	config.LoadEnv()
	r := gin.Default()

	r.Use(middleware.ErrorHandler())

	userService := userService.New()
	userHandler := user.New(userService)

	listingService := listingService.New(userService)
	listingHandler := listing.New(userService, listingService)

	api := r.Group("/public-api")
	user.RegisterRoutes(api, userHandler)
	listing.RegisterRoutes(api, listingHandler)

	r.Run(":9000")
}
