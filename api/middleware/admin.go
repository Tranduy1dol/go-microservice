package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("user_role")
		if role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			return
		}

		ctx.Next()
	}
}
