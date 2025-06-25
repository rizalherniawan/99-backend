package listing

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, handler *ListingHandler) {
	r.POST("/listings", handler.CreateListing)
	r.GET("/listings", handler.GetAllListing)
}
