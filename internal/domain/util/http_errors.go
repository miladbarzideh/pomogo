package util

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

const (
	BadRequestErrMessage     = "Bad request"
	InternalServerErrMessage = "Internal server error"
	NotFoundErrMessage       = "Not found"
)

type ResponseError struct {
	ErrStatus int
	ErrError  string
	ErrCauses interface{}
}

type ResponseErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

func (e ResponseError) Status() int {
	return e.ErrStatus
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e ResponseError) Causes() interface{} {
	return e.ErrCauses
}

func NewResponseError(status int, err string, causes interface{}) ResponseError {
	return ResponseError{
		ErrError:  err,
		ErrStatus: status,
		ErrCauses: causes,
	}
}

func NewBadRequest(causes interface{}) ResponseError {
	return NewResponseError(http.StatusBadRequest, BadRequestErrMessage, causes)
}

func NewInternalServerError(causes interface{}) ResponseError {
	return NewResponseError(http.StatusInternalServerError, InternalServerErrMessage, causes)
}

func NewNotFoundError(causes interface{}) ResponseError {
	return NewResponseError(http.StatusNotFound, NotFoundErrMessage, causes)
}

func mapErrors(err error) ResponseError {
	var resErr ResponseError
	switch {
	case errors.As(err, &resErr):
		return resErr
	case errors.Is(err, gorm.ErrRecordNotFound):
		return NewNotFoundError(err)
	case errors.Is(err, fiber.ErrUnprocessableEntity):
		return NewResponseError(fiber.ErrUnprocessableEntity.Code, fiber.ErrUnprocessableEntity.Message, err)
	default:
		return NewResponseError(http.StatusInternalServerError, err.Error(), err)
	}
}

func ParseError(err error) (int, string) {
	responseError := mapErrors(err)
	return responseError.ErrStatus, responseError.Error()
}
