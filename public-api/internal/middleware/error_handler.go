package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/common"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/exception"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		resp := common.ErrResponseDto{
			Results: false,
		}
		if len(ctx.Errors) > 0 {
			lastErr := ctx.Errors.Last().Err
			switch e := lastErr.(type) {
			case *exception.ApiError:
				resp.Errors = e.Message
				ctx.JSON(e.StatusCode, resp)
			case validator.ValidationErrors:
				resp.Errors = fmt.Sprintf("field of %s is required", e[0].Field())
				ctx.JSON(http.StatusNotAcceptable, resp)
			default:
				ctx.Status(http.StatusInternalServerError)
				resp.Errors = "internal server error"
				ctx.JSON(http.StatusInternalServerError, resp)
			}
			return
		}
	}
}
