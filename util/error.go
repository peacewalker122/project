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

//
//func PointerConvertMultiError(multiErr *MultiError) MultiError {
//	if multiErr == nil {
//		return MultiError{}
//	}
//	return *multiErr
//}

type MultiError struct {
	Errors []error
}

func (e *MultiError) Add(err error) {
	e.Errors = append(e.Errors, err)
}

func (e *MultiError) Error() string {
	var s []string
	for _, e := range e.Errors {
		s = append(s, e.Error())
	}
	return strings.Join(s, ", ")
}

func (e *MultiError) HasError() bool {
	return len(e.Errors) > 0
}
func (e *MultiError) Has(errors string) bool {
	for _, err := range e.Errors {
		if err.Error() == errors {
			return true
		}
	}
	return false
}
