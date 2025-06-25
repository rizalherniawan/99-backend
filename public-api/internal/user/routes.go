package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, handler *UserHandler) {
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUserById)
}
