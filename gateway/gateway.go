package gateway

import (
	"event-booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	authtoken := context.Request.Header.Get("Authorization")

	if authtoken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Auth token required"})
		return
	}

	userId, err := utils.VerifyToken(authtoken)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Auth token Invalid"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
