package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/common"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/exception"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		resp := common.ErrorBaseResponseDto{
			Results: false,
		}
		if len(ctx.Errors) > 0 {
			lastErr := ctx.Errors.Last().Err
			switch e := lastErr.(type) {
			case *exception.ApiError:
				ctx.Status(e.StatusCode)
				resp.Errors = e.Message
				ctx.JSON(e.StatusCode, resp)
			case validator.ValidationErrors:
				ctx.Status(http.StatusBadRequest)
				resp.Errors = fmt.Sprintf("field of %s is required", e[0].Field())
				ctx.JSON(http.StatusBadRequest, resp)
			default:
				ctx.Status(http.StatusInternalServerError)
				resp.Errors = "internal server error"
				ctx.JSON(http.StatusInternalServerError, resp)
			}
			return
		}
	}
}
