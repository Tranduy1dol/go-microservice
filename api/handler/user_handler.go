package handler

import (
	"net/http"

	"github.com/Tranduy1dol/learning-japanese/api/apperror"
	"github.com/Tranduy1dol/learning-japanese/internal/port"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo port.UserRepository
}

func NewUserHandler(userRepo port.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

// GetMe godoc
// @Summary     Get current user profile
// @Tags        users
// @Produce     json
// @Success     200 {object} domain.User
// @Failure     401 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Security    BearerAuth
// @Router      /users/me [get]
func (h *UserHandler) GetMe(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := h.userRepo.GetByID(ctx.Request.Context(), userID)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
