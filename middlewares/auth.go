package middlewares

import (
	"net/http"

	"example.com/events-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized!"})
		return
	}

	userID, isAdmin, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid auth token!"})
		return
	}

	ctx.Set("userID", userID)
	ctx.Set("isAdmin", isAdmin)

	ctx.Next()
}

func AdminCheck(ctx *gin.Context) {
	isAdmin := ctx.GetBool("isAdmin")
	if !isAdmin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "This URL requires admin status."})
		return
	}

	ctx.Next()
}
