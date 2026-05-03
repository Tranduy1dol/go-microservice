package apperror

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func FromValidationError(err error) *AppError {
	if ve, ok := errors.AsType[validator.ValidationErrors](err); ok {
		fields := make([]string, len(ve))
		for i, fe := range ve {
			fields[i] = fmt.Sprintf("%s: failed on '%s'", fe.Field(), fe.Tag())
		}

		return &AppError{
			Code:    400,
			Message: "validation failed",
			Detail:  strings.Join(fields, "; "),
		}
	}

	return BadRequest("invalid request body")
}

func Response(ctx *gin.Context, err error) {
	appErr := FromDomainError(err)
	ctx.JSON(appErr.Code, appErr)
}
