package errory

import (
	"fmt"
)

type Error struct {
	code int
	msg  string
}

const (
	defaultErrorCode = -1
	unknownErrorCode = -2
)

func (e *Error) Error() string {
	return e.msg
}

func New(message string) *Error {
	return NewCustom(defaultErrorCode, message)
}

func Newf(format string, args ...interface{}) *Error {
	return NewCustom(defaultErrorCode, fmt.Sprintf(format, args...))
}

func NewCustom(c int, message string) *Error {
	return &Error{
		code: c,
		msg:  message,
	}
}
