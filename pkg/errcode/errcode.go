package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var (
	codes = make(map[int]string)
)

func NewError(code int, message string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("錯誤 %d 已存在，請更換一個", code))
	}
	codes[code] = message
	return &Error{code: code, msg: message}
}

func (e *Error) Error() string {
	return fmt.Sprintf("錯誤:%d, 錯誤訊息:%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(detail ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range e.details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
