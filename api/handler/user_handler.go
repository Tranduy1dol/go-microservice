package handler

import (
	"net/http"

	"github.com/Tranduy1dol/learning-japanese/internal/port"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo port.UserRepository
}

func NewUserHandler(userRepo port.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) GetMe(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := h.userRepo.GetByID(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
