package errors

import "fmt"

type defaultError struct {
	code    ErrorCode
	message string
}

func (d *defaultError) SetMessage(msg string) {
	d.message = msg
}

func (d *defaultError) Message() string {
	return d.message
}

func (d *defaultError) Error() string {
	return d.message
}

func (d *defaultError) Code() ErrorCode {
	return d.code
}

func NewError(code ErrorCode, msg string) Error {
	return &defaultError{
		code:    code,
		message: msg,
	}
}

func NewDefaultError(err error) Error {
	return &defaultError{
		code:    ErrCodeInternalServer,
		message: fmt.Sprintf("%v", err),
	}
}
