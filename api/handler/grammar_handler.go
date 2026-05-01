package handler

import (
	"net/http"
	"strconv"

	"github.com/Tranduy1dol/learning-japanese/internal/usecase"
	"github.com/gin-gonic/gin"
)

type GrammarHandler struct {
	lookupSvc *usecase.LookupService
}

func NewGrammarHandler(lookupSvc *usecase.LookupService) *GrammarHandler {
	return &GrammarHandler{lookupSvc: lookupSvc}
}

func (h *GrammarHandler) GetGrammar(ctx *gin.Context) {
	id := ctx.Param("id")
	grammar, err := h.lookupSvc.GetGrammar(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "grammar not found"})
		return
	}
	ctx.JSON(http.StatusOK, grammar)
}

func (h *GrammarHandler) ListGrammar(ctx *gin.Context) {
	level, _ := strconv.Atoi(ctx.DefaultQuery("jlpt", "5"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))

	grammars, err := h.lookupSvc.ListGrammarByJLPT(ctx.Request.Context(), level, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"grammars": grammars})
}
