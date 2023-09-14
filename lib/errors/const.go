package errors

import "github.com/gofiber/fiber/v2"

var (
	ErrInternalServer   = fiber.NewError(ErrCodeInternalServer, "Internal server error")
	ErrNotFound         = fiber.NewError(ErrCodeInternalServer, "Not found")
	ErrInvalidToken     = fiber.NewError(ErrCodeInvalidJWT, "Invalid or expired access token")
	ErrPermissionDenied = fiber.NewError(ErrCodePermissionDenied, "Permission Denied")
	ErrBadParameter     = fiber.NewError(ErrCodeBadRequest, "Bad parameter")
)
