package errors

type ErrorCode int

const (
	Success                 = 0
	ErrCode2FA              = 203
	ErrCodeInvalidJWT       = 401
	ErrCodeBadRequest       = 402
	ErrCodePermissionDenied = 403
	ErrCodeNotFound         = 404
	ErrCodeInternalServer   = 500
)
