package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"success":    false,
				"error_code": 40103,
				"message":    "unauthorized",
			})
			context.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"success":    false,
				"error_code": 40103,
				"message":    "unauthorized",
			})
			context.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			context.Set("user", claims)
			context.Next()
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{
				"success":    false,
				"error_code": 40103,
				"message":    "unauthorized",
			})
			context.Abort()
		}
	}
}
