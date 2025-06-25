package service

import (
	"github.com/rizalherniawan/99-backend-test/public-api/internal/exception"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/user/model"
)

type UserService interface {
	CreateUser(name string) (*model.User, *exception.ApiError)
	GetUserById(id string) (*model.User, *exception.ApiError)
}
