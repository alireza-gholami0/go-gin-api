package middlewares

import (
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"github.com/alireza-gholami0/go-gin-api/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Missing authorization header"})
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]
		id, err := services.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.Set("x-user-id", id)
		c.Next()
		return
	}
}
