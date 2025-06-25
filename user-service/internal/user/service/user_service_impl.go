package service

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/rizalherniawan/99-backend-test/user-service/internal/exception"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/user/model"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/user/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func New(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

// Add user
func (u *UserServiceImpl) AddUser(name string) (*model.User, *exception.ApiError) {
	now := time.Now().UnixMicro()
	id, er := u.UserRepository.CreateUser(name, now, now)
	if er != nil {
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}
	return &model.User{Id: id, Name: name, CreatedAt: now, UpdatedAt: now}, nil
}

// Get all users
func (u *UserServiceImpl) GetAllUsers(size int, page int) ([]model.User, *exception.ApiError) {
	log.Println("users")
	offset := (page - 1) * size
	res, er := u.UserRepository.GetAllUsers(size, offset)
	if er != nil && !errors.Is(er, sql.ErrNoRows) {
		return []model.User{}, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}
	log.Println(res)
	return res, nil
}

// Get user by id
func (u *UserServiceImpl) GetUserById(id int) (*model.User, *exception.ApiError) {
	res, er := u.UserRepository.GetUserById(id)
	if er != nil {
		if errors.Is(er, sql.ErrNoRows) {
			return nil, exception.NewApiError("data not found", http.StatusNotFound)
		}
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}
	return res, nil
}
