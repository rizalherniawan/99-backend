package repository

import (
	"database/sql"
	"log"

	"github.com/rizalherniawan/99-backend-test/user-service/internal/user/model"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func New(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) CreateUser(name string, createdAt int64, updatedAt int64) (int64, error) {
	query := "INSERT INTO users (name, created_at, updated_at) VALUES (?, ?, ?)"
	res, er := u.db.Exec(query, name, createdAt, updatedAt)
	if er != nil {
		log.Println(er.Error())
		return 0, er
	}
	lastId, er := res.LastInsertId()
	if er != nil {
		return 0, er
	}
	return lastId, nil
}

func (u *UserRepositoryImpl) GetUserById(id int) (*model.User, error) {
	var model model.User
	query := "SELECT * FROM users WHERE id = ?"
	row := u.db.QueryRow(query, id)
	err := row.Scan(&model.Id, &model.Name, &model.CreatedAt, &model.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (u *UserRepositoryImpl) GetAllUsers(size int, offset int) ([]model.User, error) {
	var res []model.User
	query := "SELECT * FROM users ORDER BY created_at DESC LIMIT ? OFFSET ?"
	rows, er := u.db.Query(query, size, offset)
	if er != nil {
		log.Println(er)
		return res, er
	}
	for rows.Next() {
		var model model.User
		er = rows.Scan(&model.Id, &model.Name, &model.CreatedAt, &model.UpdatedAt)
		if er != nil {
			log.Println(er)
			return res, er
		}
		res = append(res, model)
	}
	return res, nil
}
