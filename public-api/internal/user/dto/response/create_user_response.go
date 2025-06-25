package response

import (
	"github.com/rizalherniawan/99-backend-test/public-api/internal/user/model"
)

type CreateUserResponse struct {
	Result bool       `json:"result"`
	Users  model.User `json:"user"`
}
