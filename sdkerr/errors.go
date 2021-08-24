package sdkerr

import (
	"errors"
	"fmt"
	"runtime"
)

const (
	TokenExpired  = "TOKEN_EXPIRED"
	CacheSetError = "CACHE_SET_ERROR"

	NoAuthCodeProvided       = "NO_AUTH_CODE_PROVIDED"
	RequestTokenError        = "REQUEST_TOKEN_ERROR"
	UnknownRequestTokenError = "UNKNOWN_REQUEST_TOKEN_ERROR"
	NoRefreshTokenProvided   = "NO_REFRESH_TOKEN_PROVIDED"

	NoApiParameterProvided = "NO_API_PARAMETER_PROVIDED"
)

type CustomError struct {
	Code    string
	Message string
	caller  string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("[%v]: code = %s message = %s", e.caller, e.Code, e.Message)
}

func NewCustomError(code, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		caller:  getCaller(),
	}
}

func NewCustomErrorf(code, format string, args ...interface{}) *CustomError {
	return NewCustomError(code, fmt.Sprintf(format, args...))
}

func FromError(err error) *CustomError {
	if err == nil {
		return nil
	}
	if se := new(CustomError); errors.As(err, &se) {
		return se
	}
	return nil
}

func getCaller() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}

	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short

	return fmt.Sprintf("%s:%d", file, line)
}
