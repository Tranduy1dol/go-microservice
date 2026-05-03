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
// @Accept      json
// @Produce     json
// @Param       req body map[string]int true "Request body containing jlpt level"
// @Success     200 {object} domain.Test
// @Failure     400 {object} map[string]string
// @Security    BearerAuth
// @Router      /tests/generate [post]
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
	ctx.JSON(http.StatusOK, test)
}
