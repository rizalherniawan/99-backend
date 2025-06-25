package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rizalherniawan/99-backend-test/public-api/config"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/exception"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/user/dto/response"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/user/model"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/util"
)

type UserServiceImpl struct {
}

func New() UserService {
	return &UserServiceImpl{}
}

func getServiceInfo() (string, string) {
	return config.GetEnv("USER_SERVICE_API_KEY"), config.GetEnv("USER_SERVICE_HOST")
}

func (u *UserServiceImpl) CreateUser(name string) (*model.User, *exception.ApiError) {
	var res response.CreateUserResponse

	// form the payload for creating user
	payload := fmt.Sprintf(`{"name":"%s"}`, name)
	body := []byte(payload)

	// get necessary information of user service
	key, host := getServiceInfo()

	// call user service for adding user
	resp, er := util.CallOtherApi("POST", host, key, bytes.NewBuffer(body), nil)
	if er != nil {
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}

	// extract the response body
	body, e := util.ExtractResponse(resp)
	if e != nil {
		return nil, e
	}
	er = json.Unmarshal(body, &res)
	if er != nil {
		log.Println("failed to unmarshal json due to: ", er.Error())
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}

	return &res.Users, nil
}

func (u *UserServiceImpl) GetUserById(id string) (*model.User, *exception.ApiError) {
	var res response.CreateUserResponse

	// get necessary information of user service
	key, host := getServiceInfo()

	// add path parameter to path
	fullPath := host + fmt.Sprintf("/%s", id)

	// call user service to get user by id
	resp, er := util.CallOtherApi("GET", fullPath, key, nil, nil)
	if er != nil {
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}

	// extract the response body
	body, e := util.ExtractResponse(resp)
	if e != nil {
		return nil, e
	}
	er = json.Unmarshal(body, &res)
	if er != nil {
		log.Println("failed to unmarshal json due to: ", er.Error())
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}
	return &res.Users, nil
}
