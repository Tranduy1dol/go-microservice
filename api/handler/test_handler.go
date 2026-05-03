package handler

import (
	"net/http"

	"github.com/Tranduy1dol/learning-japanese/api/apperror"
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
	var req struct {
		JLPT int `json:"jlpt"`
	}
	if err := ctx.BindJSON(&req); err != nil || req.JLPT < 1 || req.JLPT > 5 {
		apperror.Response(ctx, err)
		return
	}

	test, err := h.testGenSvc.GenerateTest(ctx.Request.Context(), req.JLPT)
	if err != nil {
		apperror.Response(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, test)
}
