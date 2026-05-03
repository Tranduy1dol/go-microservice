package apperror

import (
	"errors"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NotFound(detail string) *AppError {
	return &AppError{Code: 404, Message: "not found", Detail: detail}
}

func BadRequest(detail string) *AppError {
	return &AppError{Code: 400, Message: "bad request", Detail: detail}
}

func Internal(detail string) *AppError {
	return &AppError{Code: 500, Message: "internal server error", Detail: detail}
}

func Unauthorized(detail string) *AppError {
	return &AppError{Code: 401, Message: "unauthorized", Detail: detail}
}

func Forbidden(detail string) *AppError {
	return &AppError{Code: 403, Message: "forbidden", Detail: detail}
}

func Conflict(detail string) *AppError {
	return &AppError{Code: 409, Message: "already exists", Detail: detail}
}

func FromDomainError(err error) *AppError {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		return NotFound("")
	case errors.Is(err, domain.ErrInvalidInput):
		return BadRequest("")
	case errors.Is(err, domain.ErrUnauthorized):
		return Unauthorized("")
	case errors.Is(err, domain.ErrForbidden):
		return Forbidden("")
	case errors.Is(err, domain.ErrAlreadyExists):
		return Conflict("")
	default:
		return Internal("")
	}
}

func Response(ctx *gin.Context, err error) {
	appErr := FromDomainError(err)
	ctx.JSON(appErr.Code, appErr)
}
