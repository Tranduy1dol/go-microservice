package handler

import (
	"net/http"

	"github.com/Tranduy1dol/learning-japanese/api/apperror"
	"github.com/Tranduy1dol/learning-japanese/api/dto"
	"github.com/Tranduy1dol/learning-japanese/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testGenSvc *usecase.TestGeneratorService
}

func NewTestHandler(testGenSvc *usecase.TestGeneratorService) *TestHandler {
	return &TestHandler{testGenSvc: testGenSvc}
}

// GenerateTest godoc
// @Summary     Generate a new test
// @Tags        tests
// @Produce     json
// @Param       level path int true "JLPT Level"
// @Success     200 {object} dto.TestResponse
// @Failure     400 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /tests/generate/{level} [post]
func (h *TestHandler) GenerateTest(ctx *gin.Context) {
	var param dto.JLPTLevelParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	test, err := h.testGenSvc.GenerateTest(ctx.Request.Context(), param.Level)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewTestResponse(test))
}

// SubmitTest godoc
// @Summary     Submit test answers
// @Tags        tests
// @Accept      json
// @Produce     json
// @Param       id path string true "Test ID"
// @Param       req body dto.SubmitTestRequest true "Test answers"
// @Success     200 {object} dto.SubmitTestResponse
// @Failure     400 {object} apperror.AppError
// @Failure     404 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Security    BearerAuth
// @Router      /tests/{id}/submit [post]
func (h *TestHandler) SubmitTest(ctx *gin.Context) {
	var param dto.IDParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	var req dto.SubmitTestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperror.Response(ctx, apperror.FromValidationError(err))
		return
	}

	score, total, err := h.testGenSvc.SubmitTest(ctx.Request.Context(), param.ID, req.Answers)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.SubmitTestResponse{
		Score: score,
		Total: total,
	})
}
