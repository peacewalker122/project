package util

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
