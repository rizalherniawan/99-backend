package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(api *gin.RouterGroup, user *UserHandler) {
	api.POST("", user.CreateUser)
	api.GET("/:id", user.GetUserById)
	api.GET("", user.GetAllUsers)
}
