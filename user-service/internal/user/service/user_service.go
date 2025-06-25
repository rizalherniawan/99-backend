package service

import (
	"github.com/rizalherniawan/99-backend-test/user-service/internal/exception"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/user/model"
)

type UserService interface {
	GetUserById(id int) (*model.User, *exception.ApiError)
	AddUser(name string) (*model.User, *exception.ApiError)
	GetAllUsers(size int, page int) ([]model.User, *exception.ApiError)
}
