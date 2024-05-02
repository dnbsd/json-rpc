package jsonrpc

import (
	"strconv"
)

type Object map[string]any

func (p Object) get(key string) (any, bool) {
	v, ok := p[key]
	return v, ok
}

func (p Object) Number(key string) (Number, error) {
	v, ok := p.get(key)
	if !ok {
		return Number{}, NewErrorParamObjectNotFound(key)
	}
	var f float64
	switch x := v.(type) {
	case float64:
		f = x
	case float32:
		f = float64(x)
	case int64:
		f = float64(x)
	case int32:
		f = float64(x)
	case int16:
		f = float64(x)
	case int8:
		f = float64(x)
	case int:
		f = float64(x)
	case uint64:
		f = float64(x)
	case uint32:
		f = float64(x)
	case uint16:
		f = float64(x)
	case uint8:
		f = float64(x)
	case uint:
		f = float64(x)
	default:
		return Number{}, NewErrorParamObjectType(key, "number")
	}
	return Number{
		v: f,
	}, nil
}

func (p Object) String(key string) (string, error) {
	v, ok := p.get(key)
	if !ok {
		return "", NewErrorParamObjectNotFound(key)
	}
	s, ok := v.(string)
	if !ok {
		return "", NewErrorParamObjectType(key, "string")
	}
	return s, nil
}

func (p Object) Object(key string) (Object, error) {
	v, ok := p.get(key)
	if !ok {
		return nil, NewErrorParamObjectNotFound(key)
	}
	o, ok := v.(map[string]any)
	if !ok {
		return nil, NewErrorParamObjectType(key, "object")
	}
	return o, nil
}

func (p Object) Array(key string) (Array, error) {
	v, ok := p.get(key)
	if !ok {
		return nil, NewErrorParamObjectNotFound(key)
	}
	o, ok := v.([]any)
	if !ok {
		return nil, NewErrorParamObjectType(key, "array")
	}
	return o, nil
}

func (p Object) Bool(key string) (bool, error) {
	v, ok := p.get(key)
	if !ok {
		return false, NewErrorParamObjectNotFound(key)
	}
	b, ok := v.(bool)
	if !ok {
		return false, NewErrorParamObjectType(key, "bool")
	}
	return b, nil
}

type Array []any

func (p Array) get(n int) (any, bool) {
	if n >= len(p) {
		return nil, false
	}
	return p[n], true
}

func (p Array) Number(n int) (Number, error) {
	v, ok := p.get(n)
	if !ok {
		return Number{}, NewErrorParamArrayNotFound(n)
	}
	var f float64
	switch x := v.(type) {
	case float64:
		f = x
	case float32:
		f = float64(x)
	case int64:
		f = float64(x)
	case int32:
		f = float64(x)
	case int16:
		f = float64(x)
	case int8:
		f = float64(x)
	case int:
		f = float64(x)
	case uint64:
		f = float64(x)
	case uint32:
		f = float64(x)
	case uint16:
		f = float64(x)
	case uint8:
		f = float64(x)
	case uint:
		f = float64(x)
	default:
		return Number{}, NewErrorParamArrayType(n, "number")
	}
	return Number{
		v: f,
	}, nil
}

func (p Array) String(n int) (string, error) {
	v, ok := p.get(n)
	if !ok {
		return "", NewErrorParamArrayNotFound(n)
	}
	s, ok := v.(string)
	if !ok {
		return "", NewErrorParamArrayType(n, "string")
	}
	return s, nil
}

func (p Array) Object(n int) (Object, error) {
	v, ok := p.get(n)
	if !ok {
		return nil, NewErrorParamArrayNotFound(n)
	}
	o, ok := v.(map[string]any)
	if !ok {
		return nil, NewErrorParamArrayType(n, "object")
	}
	return o, nil
}

func (p Array) Array(n int) (Array, error) {
	v, ok := p.get(n)
	if !ok {
		return nil, NewErrorParamArrayNotFound(n)
	}
	o, ok := v.([]any)
	if !ok {
		return nil, NewErrorParamArrayType(n, "array")
	}
	return o, nil
}

func (p Array) Bool(n int) (bool, error) {
	v, ok := p.get(n)
	if !ok {
		return false, NewErrorParamArrayNotFound(n)
	}
	b, ok := v.(bool)
	if !ok {
		return false, NewErrorParamArrayType(n, "boolean")
	}
	return b, nil
}

type Number struct {
	v float64
}

func (n Number) Int() int {
	return int(n.v)
}

func (n Number) Uint() uint {
	return uint(n.v)
}

func (n Number) Float64() float64 {
	return n.v
}

func NewErrorParamObjectType(key, ttype string) *Error {
	code := 400001
	message := "parameter '" + key + "' is not of type " + ttype
	return NewError(code, message)
}

func NewErrorParamObjectNotFound(key string) *Error {
	code := 400002
	message := "parameter '" + key + "' not found"
	return NewError(code, message)
}

func NewErrorParamArrayType(n int, ttype string) *Error {
	code := 400003
	index := strconv.FormatInt(int64(n), 10)
	message := "parameter at position " + index + " is not of type " + ttype
	return NewError(code, message)
}

type ErrParamArrayNotFound struct {
	Index int
}

func NewErrorParamArrayNotFound(n int) *Error {
	code := 400004
	index := strconv.FormatInt(int64(n), 10)
	message := "parameter '" + index + "' not found"
	return NewError(code, message)
}
