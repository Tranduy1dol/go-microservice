package handler

import (
	"net/http"

	"github.com/Tranduy1dol/learning-japanese/api/apperror"
	"github.com/Tranduy1dol/learning-japanese/api/dto"
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
	var req dto.CreateWordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	word := req.ToDomain()
	if err := h.wordRepo.Create(ctx.Request.Context(), word); err != nil {
		apperror.Response(ctx, err)
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
	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	if err := h.wordRepo.Delete(ctx.Request.Context(), param.ID); err != nil {
		apperror.Response(ctx, err)
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
	var req dto.CreateQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	if req.CorrectIndex >= len(req.Choices) {
		apperror.Response(ctx, apperror.BadRequest("correct index out of choices"))
		return
	}

	question := req.ToDomain()
	if err := h.questionRepo.Create(ctx.Request.Context(), question); err != nil {
		apperror.Response(ctx, err)
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
	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	if err := h.questionRepo.Delete(ctx.Request.Context(), param.ID); err != nil {
		apperror.Response(ctx, err)
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
	var req dto.CreateParagraphRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	paragraph := req.ToDomain()
	if err := h.paragraphRepo.Create(ctx.Request.Context(), paragraph); err != nil {
		apperror.Response(ctx, err)
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
	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	if err := h.paragraphRepo.Delete(ctx.Request.Context(), param.ID); err != nil {
		apperror.Response(ctx, err)
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
	var req dto.CreateGrammarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, err)
		return
	}

	grammar := req.ToDomain()
	if err := h.grammarRepo.Create(ctx.Request.Context(), grammar); err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"grammar": req})
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
	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	if err := h.grammarRepo.Delete(ctx.Request.Context(), param.ID); err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
