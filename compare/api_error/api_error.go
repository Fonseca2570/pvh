package api_error

import "fmt"

//swagger:model Error
type ApiError struct {
	Msg string `json:"msg"`
}

func (e *ApiError) Error() string{
	fmt.Println(e.Msg)
	return e.Msg
}

func NewError(msg string) error {
	return &ApiError{Msg: msg}
}