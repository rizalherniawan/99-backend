package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizalherniawan/99-backend-test/user-service/config"
	"github.com/rizalherniawan/99-backend-test/user-service/internal/common"
)

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedKey := config.GetEnv("INTERNAL_API_KEY")
		providedKey := c.GetHeader("X-API-KEY")

		if expectedKey == "" || providedKey != expectedKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorBaseResponseDto{
				Results: false,
				Errors:  "Unauthorized",
			})
			return
		}

		// Proceed to the next middleware or handler
		c.Next()
	}
}
