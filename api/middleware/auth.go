package middleware

import (
	"net/http"
	"strings"

	"github.com/Tranduy1dol/learning-japanese/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSvc *auth.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
			return
		}

		claims, err := jwtSvc.ValidateToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("user_role", claims.Role)

		ctx.Next()
	}
}
