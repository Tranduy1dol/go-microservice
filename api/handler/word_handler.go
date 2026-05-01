package handler

import (
	"net/http"
	"strconv"

	"github.com/Tranduy1dol/learning-japanese/internal/usecase"
	"github.com/gin-gonic/gin"
)

type WordHandler struct {
	lookupSvc *usecase.LookupService
}

func NewWordHandler(lookupSvc *usecase.LookupService) *WordHandler {
	return &WordHandler{lookupSvc: lookupSvc}
}

func (h *WordHandler) GetWord(ctx *gin.Context) {
	id := ctx.Param("id")
	word, err := h.lookupSvc.GetWord(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "word not found"})
		return
	}

	ctx.JSON(http.StatusOK, word)
}

func (h *WordHandler) SearchWords(ctx *gin.Context) {
	query := ctx.Query("q")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query required"})
		return
	}

	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	words, err := h.lookupSvc.SearchWord(ctx.Request.Context(), query, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"words": words})
}

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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"words":  words,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}
