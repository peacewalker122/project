package util

import "strings"

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() (int, string) {
	return e.Code, e.Message
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

type MultiError struct {
	Errors []*Error
}

func (e *MultiError) Add(err error, code ...int) {
	e.Errors = append(e.Errors, NewError(code[0], err.Error()))
}

func (e *MultiError) Error() string {
	var s []string
	for _, e := range e.Errors {
		s = append(s, e.Message)
	}
	return strings.Join(s, ", ")
}

func (e *MultiError) HasError() bool {
	return len(e.Errors) > 0
}
func (e *MultiError) Has(errors string) bool {
	for _, err := range e.Errors {
		if err.Message == errors {
			return true
		}
	}
	return false
}
