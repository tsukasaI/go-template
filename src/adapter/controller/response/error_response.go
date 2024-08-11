package response

import (
	customErrors "my-go-template/src/core/errors"
	"time"
)

type (
	ErrorResponse struct {
		RequestID string `json:"request_id"`
		Timestamp string `json:"timestamp"`
		Error     Error  `json:"error"`
	}
	Error struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}
)

func NewErrorResponse(requestID string, err customErrors.ErrInterface) *ErrorResponse {
	return &ErrorResponse{
		RequestID: requestID,
		Timestamp: time.Now().Format("2006-01-02T15:04:05+09:00"),
		Error:     *newError(err),
	}
}

func NewEchoErrorResponse(requestID string, err customErrors.ErrInterface) (int, any) {
	return err.GetStatusCode(), &ErrorResponse{
		RequestID: requestID,
		Timestamp: time.Now().Format("2006-01-02T15:04:05+09:00"),
		Error:     *newError(err),
	}
}

func newError(err customErrors.ErrInterface) *Error {
	return &Error{
		Message: err.GetMessage(),
		Code:    err.GetCode(),
	}
}
