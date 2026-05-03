package handler

import (
	"net/http"
	"strconv"

	"github.com/Tranduy1dol/learning-japanese/api/apperror"
	"github.com/Tranduy1dol/learning-japanese/internal/usecase"
	"github.com/gin-gonic/gin"
)

type GrammarHandler struct {
	lookupSvc *usecase.LookupService
}

func NewGrammarHandler(lookupSvc *usecase.LookupService) *GrammarHandler {
	return &GrammarHandler{lookupSvc: lookupSvc}
}

// GetGrammar godoc
// @Summary     Get grammar by ID
// @Tags        grammar
// @Produce     json
// @Param       id path string true "Grammar ID"
// @Success     200 {object} domain.Grammar
// @Failure     404 {object} map[string]string
// @Security    BearerAuth
// @Router      /grammar/{id} [get]
func (h *GrammarHandler) GetGrammar(ctx *gin.Context) {
	id := ctx.Param("id")
	grammar, err := h.lookupSvc.GetGrammar(ctx.Request.Context(), id)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, grammar)
}

// ListGrammar godoc
// @Summary     List grammar by JLPT level
// @Tags        grammar
// @Produce     json
// @Param       jlpt query int false "JLPT Level" default(5)
// @Param       limit query int false "Limit" default(50)
// @Success     200 {object} map[string]interface{}
// @Security    BearerAuth
// @Router      /grammar [get]
func (h *GrammarHandler) ListGrammar(ctx *gin.Context) {
	level, _ := strconv.Atoi(ctx.DefaultQuery("jlpt", "5"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))

	grammars, err := h.lookupSvc.ListGrammarByJLPT(ctx.Request.Context(), level, limit)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"grammars": grammars})
}
