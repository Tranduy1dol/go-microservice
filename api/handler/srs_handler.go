package handler

import (
	"net/http"

	"github.com/Tranduy1dol/learning-japanese/api/apperror"
	"github.com/Tranduy1dol/learning-japanese/api/dto"
	"github.com/Tranduy1dol/learning-japanese/internal/usecase"
	"github.com/gin-gonic/gin"
)

type SRSHandler struct {
	srsSvc *usecase.SRSService
}

func NewSRSHandler(srsSvc *usecase.SRSService) *SRSHandler {
	return &SRSHandler{srsSvc: srsSvc}
}

// AddWordToDeck godoc
// @Summary     Add a word to user's SRS deck
// @Tags        srs
// @Accept      json
// @Produce     json
// @Param       req body dto.AddWordToDeckRequest true "Word ID"
// @Success     201 {object} dto.SRSCardResponse
// @Failure     400 {object} apperror.AppError
// @Failure     409 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /srs/deck [post]
func (h *SRSHandler) AddWordToDeck(ctx *gin.Context) {
	userID, err := dto.UserIDFromContext(ctx)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	var req dto.AddWordToDeckRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	card, err := h.srsSvc.AddWordToDeck(ctx.Request.Context(), userID, req.WordID)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewSRSCardResponse(card))
}

// GetDueCards godoc
// @Summary     Get flashcards due for review
// @Tags        srs
// @Produce     json
// @Param       limit query int false "Limit" default(20)
// @Success     200 {array} dto.SRSCardResponse
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /srs/due [get]
func (h *SRSHandler) GetDueCards(ctx *gin.Context) {
	userID, err := dto.UserIDFromContext(ctx)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	var page dto.PaginationQuery
	if err := ctx.ShouldBindQuery(&page); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	cards, err := h.srsSvc.GetDueCards(ctx.Request.Context(), userID, page.Limit)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewSRSCardListResponse(cards))
}

// ReviewCard godoc
// @Summary     Submit a flashcard review
// @Tags        srs
// @Accept      json
// @Produce     json
// @Param       id path string true "Flashcard ID"
// @Param       req body dto.ReviewCardRequest true "Review quality (0-5)"
// @Success     200 {object} dto.SRSCardResponse
// @Failure     400 {object} apperror.AppError
// @Failure     404 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /srs/review/{id} [post]
func (h *SRSHandler) ReviewCard(ctx *gin.Context) {
	userID, err := dto.UserIDFromContext(ctx)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	var req dto.ReviewCardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	card, err := h.srsSvc.ReviewCard(ctx, userID, param.ID, req.Quality)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewSRSCardResponse(card))
}
