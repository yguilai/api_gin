package error

import "errors"

const (
	SUCCESS    = 0
	ERROR_AUTH = 400 + iota
	ERROR_AUTH_TOKEN
	ERROR_AUTH_CHECK_TOKEN_FAIL
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	INVALID_PARAMS
)

func GetMsg(code int) string {
	Msg := ""
	switch code {
	case ERROR_AUTH:
		Msg = "Unauthorized access to this API"
	case ERROR_AUTH_TOKEN:
		Msg = "Generate token has failed"
	case ERROR_AUTH_CHECK_TOKEN_FAIL:
		Msg = "Token validation failed"
	case ERROR_AUTH_CHECK_TOKEN_TIMEOUT:
		Msg = "Token has expired"
	case INVALID_PARAMS:
		Msg = "Invalid param"
	}
	err := errors.New(Msg)
	return err.Error()
}
