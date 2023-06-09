package middlewares

import (
	authentification"grpc-auth-micro/authentifiacation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString, err := context.Cookie("UserJWT")
		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Request does not contain an access token",
			})
			context.Abort()
			return
		}
		err = authentification.ValidateToken(tokenString)
		context.Set("user", authentification.P)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
