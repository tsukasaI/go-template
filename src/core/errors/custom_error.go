package errors

import "log"

const (
	NotFoundErrCode       = "NotFound"
	BadRequestErrCode     = "BadRequest"
	InternalServerErrCode = "InternalServer"
)

type (
	Err struct {
		Error      error  `json:"-"`
		StatusCode int    `json:"-"`
		Code       string `json:"code"`
		Message    string `json:"message"`
	}
	ErrInterface interface {
		GetError() error
		GetStatusCode() int
		GetCode() string
		GetMessage() string
	}
)

func (e *Err) GetError() error {
	return e.Error
}

func (e *Err) GetStatusCode() int {
	return e.StatusCode
}

func (e *Err) GetCode() string {
	return e.Code
}

func (e *Err) GetMessage() string {
	return e.Message
}

func NewErr(err error, statusCode int, code, message string) ErrInterface {
	log.Printf("%v\n", err)
	return &Err{
		Error:      err,
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}
