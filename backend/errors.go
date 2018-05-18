package backend

import (
	"fmt"
	"strings"
)

const (
	ParamErrorTypeMissing = 1
	ParamErrorTypeInvalid = 2
)

type ParamError struct {
	Msg    string
	Type   int
	Fields []string
}

func (data *ParamError) Code() string {

	var typ string

	switch data.Type {
	case ParamErrorTypeMissing:
		typ = "missing-parameter"
	case ParamErrorTypeInvalid:
		typ = "invalid-parameter"
	}

	return fmt.Sprintf("%s:%s", typ, strings.Join(data.Fields, ","))
}

func (data *ParamError) Error() string {

	return fmt.Sprintf("%s:%s", data.Code(), data.Msg)
}

func NewParamError(typ int, msg string, fields ...string) error {

	return &ParamError{Msg: msg, Type: typ, Fields: fields}
}

type LogicError struct {
	Code string
	Msg  string
}

func (data *LogicError) Error() string {

	return fmt.Sprintf("%s:%s", data.Code, data.Msg)
}
