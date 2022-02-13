package domain

import "net/http"

type CustomError struct {
	msg  string
	Code int
}

func NewError(msg string, code int) *CustomError {
	return &CustomError{
		msg:  msg,
		Code: code,
	}
}

func (c *CustomError) Error() string {
	return c.msg
}

var (
	ErrorInternalServerError = NewError("internal server error", http.StatusInternalServerError)
	ErrorBadRequest          = NewError("bad query input", http.StatusBadRequest)
	ErrorUserNotExist        = NewError("user does not exist", http.StatusBadRequest)
	ErrorUserIsExist         = NewError("nickname already registered or email already registered", http.StatusConflict)
)
