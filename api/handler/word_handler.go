package handler

import (
	"net/http"

	"github.com/Tranduy1dol/learning-japanese/api/apperror"
	"github.com/Tranduy1dol/learning-japanese/api/dto"
	"github.com/Tranduy1dol/learning-japanese/internal/usecase"
	"github.com/gin-gonic/gin"
)

type WordHandler struct {
	lookupSvc *usecase.LookupService
}

func NewWordHandler(lookupSvc *usecase.LookupService) *WordHandler {
	return &WordHandler{lookupSvc: lookupSvc}
}

// GetWord godoc
// @Summary     Get word by ID
// @Tags        words
// @Produce     json
// @Param       id path string true "Word ID"
// @Success     200 {object} domain.Word
// @Failure     404 {object} map[string]string
// @Security    BearerAuth
// @Router      /words/{id} [get]
func (h *WordHandler) GetWord(ctx *gin.Context) {
	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	word, err := h.lookupSvc.GetWord(ctx.Request.Context(), param.ID)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewWordResponse(word))
}

// SearchWords godoc
// @Summary     Search words
// @Tags        words
// @Produce     json
// @Param       q     query string true  "Search query"
// @Param       limit query int    false "Limit" default(20)
// @Success     200 {object} map[string]interface{}
// @Security    BearerAuth
// @Router      /words/search [get]
func (h *WordHandler) SearchWords(ctx *gin.Context) {
	var query dto.SearchWordQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	words, err := h.lookupSvc.SearchWord(ctx.Request.Context(), query.Q, query.Limit)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewWordListResponse(words))
}

// BrowseWordsByJLPT godoc
// @Summary     Browse words by JLPT level
// @Tags        words
// @Produce     json
// @Param       level path int true "JLPT Level (1-5)"
// @Param       limit query int false "Limit" default(50)
// @Param       offset query int false "Offset" default(0)
// @Success     200 {object} map[string]interface{}
// @Security    BearerAuth
// @Router      /words/jlpt/{level} [get]
func (h *WordHandler) BrowseWordsByJLPT(ctx *gin.Context) {
	var param dto.JLPTLevelParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	var page dto.PaginationQuery
	if err := ctx.ShouldBindQuery(&page); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	words, total, err := h.lookupSvc.BrowseWordByJLPT(ctx.Request.Context(), param.Level, page.Limit, page.Offset)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"words":  dto.NewWordListResponse(words),
		"total":  total,
		"limit":  page.Limit,
		"offset": page.Offset,
	})
}
