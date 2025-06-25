package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/dto/request"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/exception"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/user/service"
)

type UserHandler struct {
	userService service.UserService
}

func New(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) GetAllUsers(r *gin.Context) {
	page := r.DefaultQuery("page_num", "1")
	size := r.DefaultQuery("page_size", "10")

	pageInt, er := strconv.Atoi(page)
	if er != nil {
		log.Println(er.Error())
		r.Status(http.StatusInternalServerError)
		r.Error(exception.NewApiError("internal server error", http.StatusInternalServerError))
		r.Abort()
		return
	}

	sizeInt, er := strconv.Atoi(size)
	if er != nil {
		log.Println(er.Error())
		r.Status(http.StatusInternalServerError)
		r.Error(exception.NewApiError("internal server error", http.StatusInternalServerError))
		r.Abort()
		return
	}

	res, e := u.userService.GetAllUsers(sizeInt, pageInt)
	if e != nil {
		r.Status(e.StatusCode)
		r.Error(e)
		r.Abort()
		return
	}

	r.JSON(http.StatusOK, map[string]interface{}{
		"result": true,
		"users":  res,
	})
}

func (u *UserHandler) CreateUser(r *gin.Context) {
	req := request.CreateUserRequest{}
	e := r.ShouldBindJSON(&req)
	if e != nil {
		r.Error(e)
		r.Abort()
		return
	}

	res, er := u.userService.AddUser(req.Name)
	if er != nil {
		r.Error(er)
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
	idInt, er := strconv.Atoi(id)
	if er != nil {
		log.Println(er.Error())
		r.Error(er)
		r.Abort()
		return
	}

	res, e := u.userService.GetUserById(idInt)
	if e != nil {
		r.Error(e)
		r.Abort()
		return
	}

	r.JSON(http.StatusOK, map[string]interface{}{
		"result": true,
		"user":   res,
	})

}
