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
		return Number{}, &ErrParamObjectNotFound{
			Key: key,
		}
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
		return Number{}, &ErrParamObjectType{
			Key:  key,
			Type: "number",
		}
	}
	return Number{
		v: f,
	}, nil
}

func (p Object) String(key string) (string, error) {
	v, ok := p.get(key)
	if !ok {
		return "", &ErrParamObjectNotFound{
			Key: key,
		}
	}
	s, ok := v.(string)
	if !ok {
		return "", &ErrParamObjectType{
			Key:  key,
			Type: "string",
		}
	}
	return s, nil
}

func (p Object) Object(key string) (Object, error) {
	v, ok := p.get(key)
	if !ok {
		return nil, &ErrParamObjectNotFound{
			Key: key,
		}
	}
	o, ok := v.(map[string]any)
	if !ok {
		return nil, &ErrParamObjectType{
			Key:  key,
			Type: "object",
		}
	}
	return o, nil
}

func (p Object) Array(key string) (Array, error) {
	v, ok := p.get(key)
	if !ok {
		return nil, &ErrParamObjectNotFound{
			Key: key,
		}
	}
	o, ok := v.([]any)
	if !ok {
		return nil, &ErrParamObjectType{
			Key:  key,
			Type: "array",
		}
	}
	return o, nil
}

func (p Object) Bool(key string) (bool, error) {
	v, ok := p.get(key)
	if !ok {
		return false, &ErrParamObjectNotFound{
			Key: key,
		}
	}
	b, ok := v.(bool)
	if !ok {
		return false, &ErrParamObjectType{
			Key:  key,
			Type: "bool",
		}
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
		return Number{}, &ErrParamArrayNotFound{
			Index: n,
		}
	}
	f, ok := v.(float64)
	if !ok {
		return Number{}, &ErrParamArrayType{
			Index: n,
			Type:  "number",
		}
	}
	return Number{
		v: f,
	}, nil
}

func (p Array) String(n int) (string, error) {
	v, ok := p.get(n)
	if !ok {
		return "", &ErrParamArrayNotFound{
			Index: n,
		}
	}
	s, ok := v.(string)
	if !ok {
		return "", &ErrParamArrayType{
			Index: n,
			Type:  "string",
		}
	}
	return s, nil
}

func (p Array) Object(n int) (Object, error) {
	v, ok := p.get(n)
	if !ok {
		return nil, &ErrParamArrayNotFound{
			Index: n,
		}
	}
	o, ok := v.(map[string]any)
	if !ok {
		return nil, &ErrParamArrayType{
			Index: n,
			Type:  "object",
		}
	}
	return o, nil
}

func (p Array) Array(n int) (Array, error) {
	v, ok := p.get(n)
	if !ok {
		return nil, &ErrParamArrayNotFound{
			Index: n,
		}
	}
	o, ok := v.([]any)
	if !ok {
		return nil, &ErrParamArrayType{
			Index: n,
			Type:  "array",
		}
	}
	return o, nil
}

func (p Array) Bool(n int) (bool, error) {
	v, ok := p.get(n)
	if !ok {
		return false, &ErrParamArrayNotFound{
			Index: n,
		}
	}
	b, ok := v.(bool)
	if !ok {
		return false, &ErrParamArrayType{
			Index: n,
			Type:  "boolean",
		}
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

type ErrParamObjectType struct {
	Key  string
	Type string
}

func (p *ErrParamObjectType) Error() string {
	return "parameter '" + p.Key + "' is not of type " + p.Type
}

type ErrParamObjectNotFound struct {
	Key string
}

func (p *ErrParamObjectNotFound) Error() string {
	return "parameter '" + p.Key + "' not found"
}

type ErrParamArrayType struct {
	Index int
	Type  string
}

func (p *ErrParamArrayType) Error() string {
	index := strconv.FormatInt(int64(p.Index), 10)
	return "parameter at position " + index + " is not of type " + p.Type
}

type ErrParamArrayNotFound struct {
	Index int
}

func (p *ErrParamArrayNotFound) Error() string {
	index := strconv.FormatInt(int64(p.Index), 10)
	return "parameter '" + index + "' not found"
}
