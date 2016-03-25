package info

import "encoding"

type Container interface {
	Get(domPath string) interface{}
	Set(domPath string, value interface{} error
	TplString(template string) string

	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
