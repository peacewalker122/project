package util

import "github.com/labstack/gommon/log"

type Error struct {
	Err map[bool]map[string]string
}

func (e *Error) Error(key string) string {
	return e.Err[true][key]
}

func (e *Error) Important(Error, key string) {
	e.Err[true][key] = Error
}

func (e *Error) Unimportant(Error, key string) {
	log.Error(Error)
	e.Err[false][key] = Error
}

func (e *Error) Errors() map[string]string {
	return e.Err[true]
}

func (e *Error) HasUnimportantError() bool {
	return len(e.Err[false]) > 0
}

func (e *Error) HasError() bool {
	return len(e.Err[true]) > 0
}
