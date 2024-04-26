package jsonrpc

import (
	"context"
	"errors"
)

type Context struct {
	ctx    context.Context
	params any
}

func (c *Context) Context() context.Context {
	return c.ctx
}

func (c *Context) Result(result any) (any, error) {
	return result, nil
}

func (c *Context) Error(err error) (any, error) {
	return nil, err
}

func (c *Context) ParamsObject() (Object, error) {
	switch v := c.params.(type) {
	case map[string]any:
		return v, nil
	case Object:
		return v, nil
	default:
		return nil, errors.New("method parameters must be an object")
	}
}

func (c *Context) ParamsArray() (Array, error) {
	switch v := c.params.(type) {
	case []any:
		return v, nil
	case Array:
		return v, nil
	default:
		return nil, errors.New("method parameters must be an array")
	}
}
