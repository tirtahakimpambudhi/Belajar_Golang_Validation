package exception

import "fmt"

type ErrorMessage struct {
	Field string
	Value interface{}
	Tag   string
}

func NewErrorMessage(field string, value interface{}, tag string) *ErrorMessage {
	return &ErrorMessage{Field: field, Value: value, Tag: tag}
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("Error : Field '%v' with Value '%v' because Tag '%v'", e.Field, e.Value, e.Tag)
}
