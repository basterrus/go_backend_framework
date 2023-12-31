package rest

import (
	"encoding/json"
	"fmt"
)

var (
	ErrNotFound   = NewAppError("not found", "NS-000003", "")
	ErrBadRequest = NewAppError("bad request", "NS-000002", "some thing wrong with user data")
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func NewAppError(message, code, developerMessage string) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Code:             code,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return bytes
}

//func BadRequestError(message string) *AppError {
//	return ErrBadRequest
//}

func systemError(developerMessage string) *AppError {
	return NewAppError("system error", "NS-000001", developerMessage)
}
