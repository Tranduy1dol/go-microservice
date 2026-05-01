package handler

import (
	"net/http"

	"github.com/Tranduy1dol/learning-japanese/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testGenSvc *usecase.TestGeneratorService
}

func NewTestHandler(testGenSvc *usecase.TestGeneratorService) *TestHandler {
	return &TestHandler{testGenSvc: testGenSvc}
}

func (h *TestHandler) GenerateTest(ctx *gin.Context) {
	var req struct {
		JLPT int `json:"jlpt"`
	}
	if err := ctx.BindJSON(&req); err != nil || req.JLPT < 1 || req.JLPT > 5 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid jlpt level"})
		return
	}

	test, err := h.testGenSvc.GenerateTest(ctx.Request.Context(), req.JLPT)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, test)
}
