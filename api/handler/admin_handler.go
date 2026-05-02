package handler

import (
	"net/http"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
	"github.com/Tranduy1dol/learning-japanese/internal/port"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	wordRepo      port.DictionaryRepository
	questionRepo  port.QuestionRepository
	paragraphRepo port.ParagraphRepository
	grammarRepo   port.GrammarRepository
}

func NewAdminHandler(
	wordRepo port.DictionaryRepository,
	questionRepo port.QuestionRepository,
	paragraphRepo port.ParagraphRepository,
	grammarRepo port.GrammarRepository,
) *AdminHandler {
	return &AdminHandler{
		wordRepo:      wordRepo,
		questionRepo:  questionRepo,
		paragraphRepo: paragraphRepo,
		grammarRepo:   grammarRepo,
	}
}

// CreateWord godoc
// @Summary     Create a word
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       word body domain.Word true "Word object"
// @Success     201 {object} map[string]interface{}
// @Security    BearerAuth
// @Router      /admin/words [post]
func (h *AdminHandler) CreateWord(ctx *gin.Context) {
	var word domain.Word
	if err := ctx.BindJSON(&word); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid word"})
		return
	}

	if err := h.wordRepo.Create(ctx.Request.Context(), &word); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"word": word})
}

// DeleteWord godoc
// @Summary     Delete a word
// @Tags        admin
// @Produce     json
// @Param       id path string true "Word ID"
// @Success     200 {object} map[string]bool
// @Security    BearerAuth
// @Router      /admin/words/{id} [delete]
func (h *AdminHandler) DeleteWord(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.wordRepo.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

// CreateQuestion godoc
// @Summary     Create a question
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       question body domain.Question true "Question object"
// @Success     201 {object} map[string]interface{}
// @Security    BearerAuth
// @Router      /admin/questions [post]
func (h *AdminHandler) CreateQuestion(ctx *gin.Context) {
	var question domain.Question
	if err := ctx.BindJSON(&question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid question"})
		return
	}

	if err := h.questionRepo.Create(ctx.Request.Context(), &question); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"question": question})
}

// DeleteQuestion godoc
// @Summary     Delete a question
// @Tags        admin
// @Produce     json
// @Param       id path string true "Question ID"
// @Success     200 {object} map[string]bool
// @Security    BearerAuth
// @Router      /admin/questions/{id} [delete]
func (h *AdminHandler) DeleteQuestion(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.questionRepo.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

// CreateParagraph godoc
// @Summary     Create a paragraph
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       paragraph body domain.Paragraph true "Paragraph object"
// @Success     201 {object} map[string]interface{}
// @Security    BearerAuth
// @Router      /admin/paragraphs [post]
func (h *AdminHandler) CreateParagraph(ctx *gin.Context) {
	var paragraph domain.Paragraph
	if err := ctx.BindJSON(&paragraph); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid paragraph"})
		return
	}

	if err := h.paragraphRepo.Create(ctx.Request.Context(), &paragraph); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"paragraph": paragraph})
}

// DeleteParagraph godoc
// @Summary     Delete a paragraph
// @Tags        admin
// @Produce     json
// @Param       id path string true "Paragraph ID"
// @Success     200 {object} map[string]bool
// @Security    BearerAuth
// @Router      /admin/paragraphs/{id} [delete]
func (h *AdminHandler) DeleteParagraph(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.paragraphRepo.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

// CreateGrammar godoc
// @Summary     Create a grammar entry
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       grammar body domain.Grammar true "Grammar object"
// @Success     201 {object} map[string]interface{}
// @Security    BearerAuth
// @Router      /admin/grammars [post]
func (h *AdminHandler) CreateGrammar(ctx *gin.Context) {
	var grammar domain.Grammar
	if err := ctx.BindJSON(&grammar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid grammar"})
		return
	}

	if err := h.grammarRepo.Create(ctx.Request.Context(), &grammar); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"grammar": grammar})
}

// DeleteGrammar godoc
// @Summary     Delete a grammar entry
// @Tags        admin
// @Produce     json
// @Param       id path string true "Grammar ID"
// @Success     200 {object} map[string]bool
// @Security    BearerAuth
// @Router      /admin/grammars/{id} [delete]
func (h *AdminHandler) DeleteGrammar(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.grammarRepo.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
