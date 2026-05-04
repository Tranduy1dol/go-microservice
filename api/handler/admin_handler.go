package handler

import (
	"net/http"

	"github.com/Tranduy1dol/kotoba-press-core/api/apperror"
	"github.com/Tranduy1dol/kotoba-press-core/api/dto"
	"github.com/Tranduy1dol/kotoba-press-core/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminSvc *usecase.AdminService
}

func NewAdminHandler(adminSvc *usecase.AdminService) *AdminHandler {
	return &AdminHandler{
		adminSvc: adminSvc,
	}
}

// CreateWord godoc
// @Summary     Create a word
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       word body dto.CreateWordRequest true "Word object"
// @Success     201 {object} dto.WordResponse
// @Failure     400 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /admin/words [post]
func (h *AdminHandler) CreateWord(ctx *gin.Context) {
	var req dto.CreateWordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	word := req.ToDomain()
	if err := h.adminSvc.CreateWord(ctx.Request.Context(), word); err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewWordResponse(word))
}

// DeleteWord godoc
// @Summary     Delete a word
// @Tags        admin
// @Produce     json
// @Param       id path string true "Word ID"
// @Success     200 {object} map[string]bool
// @Failure     400 {object} apperror.AppError
// @Failure     404 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /admin/words/{id} [delete]
func (h *AdminHandler) DeleteWord(ctx *gin.Context) {
	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	if err := h.adminSvc.DeleteWord(ctx.Request.Context(), param.ID); err != nil {
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
// @Param       question body dto.CreateQuestionRequest true "Question object"
// @Success     201 {object} dto.QuestionWithAnswerResponse
// @Failure     400 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /admin/questions [post]
func (h *AdminHandler) CreateQuestion(ctx *gin.Context) {
	var req dto.CreateQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	question := req.ToDomain()
	if err := h.adminSvc.CreateQuestion(ctx.Request.Context(), question); err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewQuestionWithAnswerResponse(question))
}

// DeleteQuestion godoc
// @Summary     Delete a question
// @Tags        admin
// @Produce     json
// @Param       id path string true "Question ID"
// @Success     200 {object} map[string]bool
// @Failure     400 {object} apperror.AppError
// @Failure     404 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /admin/questions/{id} [delete]
func (h *AdminHandler) DeleteQuestion(ctx *gin.Context) {
	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	if err := h.adminSvc.DeleteQuestion(ctx.Request.Context(), param.ID); err != nil {
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
// @Param       paragraph body dto.CreateParagraphRequest true "Paragraph object"
// @Success     201 {object} dto.ParagraphResponse
// @Failure     400 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /admin/paragraphs [post]
func (h *AdminHandler) CreateParagraph(ctx *gin.Context) {
	var req dto.CreateParagraphRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	paragraph := req.ToDomain()
	if err := h.adminSvc.CreateParagraph(ctx.Request.Context(), paragraph); err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewParagraphResponse(paragraph))
}

// DeleteParagraph godoc
// @Summary     Delete a paragraph
// @Tags        admin
// @Produce     json
// @Param       id path string true "Paragraph ID"
// @Success     200 {object} map[string]bool
// @Failure     400 {object} apperror.AppError
// @Failure     404 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /admin/paragraphs/{id} [delete]
func (h *AdminHandler) DeleteParagraph(ctx *gin.Context) {
	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	if err := h.adminSvc.DeleteParagraph(ctx.Request.Context(), param.ID); err != nil {
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
// @Param       grammar body dto.CreateGrammarRequest true "Grammar object"
// @Success     201 {object} dto.GrammarResponse
// @Failure     400 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /admin/grammars [post]
func (h *AdminHandler) CreateGrammar(ctx *gin.Context) {
	var req dto.CreateGrammarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, err)
		return
	}

	grammar := req.ToDomain()
	if err := h.adminSvc.CreateGrammar(ctx.Request.Context(), grammar); err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewGrammarResponse(grammar))
}

// DeleteGrammar godoc
// @Summary     Delete a grammar entry
// @Tags        admin
// @Produce     json
// @Param       id path string true "Grammar ID"
// @Success     200 {object} map[string]bool
// @Failure     400 {object} apperror.AppError
// @Failure     404 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /admin/grammars/{id} [delete]
func (h *AdminHandler) DeleteGrammar(ctx *gin.Context) {
	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	if err := h.adminSvc.DeleteGrammar(ctx.Request.Context(), param.ID); err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
