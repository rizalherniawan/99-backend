package repository

import "github.com/rizalherniawan/99-backend-test/user-service/internal/user/model"

type UserRepository interface {
	CreateUser(name string, createdAt int64, updatedAt int64) (int64, error)
	GetUserById(id int) (*model.User, error)
	GetAllUsers(size int, offset int) ([]model.User, error)
}
