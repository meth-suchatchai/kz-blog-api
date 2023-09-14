package errors

import "github.com/gofiber/fiber/v2"

func NewDefaultFiberError(err error) *fiber.Error {
	return &fiber.Error{
		Code:    ErrCodeInternalServer,
		Message: err.Error(),
	}
}

func NewDefaultFiberMessageError(msg string) *fiber.Error {
	return &fiber.Error{
		Code:    ErrCodeInternalServer,
		Message: msg,
	}
}
