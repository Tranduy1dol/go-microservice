package handler

import (
	"net/http"
	"strconv"

	"github.com/Tranduy1dol/learning-japanese/api/apperror"
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
	id := ctx.Param("id")
	word, err := h.lookupSvc.GetWord(ctx.Request.Context(), id)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, word)
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
	query := ctx.Query("q")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query required"})
		return
	}

	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	words, err := h.lookupSvc.SearchWord(ctx.Request.Context(), query, limit)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"words": words})
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
	level, err := strconv.Atoi(ctx.Param("level"))
	if err != nil || level < 1 || level > 5 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid jlpt level"})
		return
	}

	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	words, total, err := h.lookupSvc.BrowseWordByJLPT(ctx.Request.Context(), level, limit, offset)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"words":  words,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}
