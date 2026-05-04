package dto

import (
	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"github.com/gin-gonic/gin"
)

type PaginationQuery struct {
	Limit  int `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset int `form:"offset" binding:"omitempty,min=0"`
}

type JLPTLevelParam struct {
	Level int `uri:"level" binding:"required,min=1,max=5"`
}

type IDParam struct {
	ID string `uri:"id" binding:"required"`
}

func UserIDFromContext(ctx *gin.Context) (string, error) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		return "", domain.ErrUnauthorized
	}

	return userID, nil
}
