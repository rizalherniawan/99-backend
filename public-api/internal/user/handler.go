package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/user/dto/request"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/user/service"
)

type UserHandler struct {
	userService service.UserService
}

func New(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) CreateUser(r *gin.Context) {
	payload := request.CreateUser{}
	er := r.ShouldBindJSON(&payload)
	if er != nil {
		r.Status(http.StatusNotAcceptable)
		r.Error(er)
		r.Abort()
		return
	}

	res, e := u.userService.CreateUser(payload.Name)
	if e != nil {
		r.Status(e.StatusCode)
		r.Error(e)
		r.Abort()
		return
	}

	r.JSON(http.StatusCreated, map[string]interface{}{
		"result": true,
		"user":   res,
	})
}

func (u *UserHandler) GetUserById(r *gin.Context) {
	id := r.Param("id")
	res, e := u.userService.GetUserById(id)
	if e != nil {
		r.Status(e.StatusCode)
		r.Error(e)
		r.Abort()
		return
	}

	r.JSON(http.StatusCreated, map[string]interface{}{
		"result": true,
		"user":   res,
	})
}
