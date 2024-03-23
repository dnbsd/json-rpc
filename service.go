package jsonrpc

import (
	"context"
)

// Service is a collection of registered modules and their methods.
type Service struct {
	methods map[string]Method
}

// Call calls a registered method using the provided request data. The method does not check if a request is a notification
// and always returns a response. It's callers responsibility to validate request data. If an unregistered method id called,
// the method returns ErrMethodNotFound. Otherwise, response data is populated by values returned from a method.
func (m *Service) Call(ctx context.Context, req Request) Response {
	method, ok := m.methods[req.Method]
	if !ok {
		return NewErrorResponse(req.ID, ErrMethodNotFound)
	}

	c := &Context{
		ctx:    ctx,
		params: req.Params,
	}
	result, err := method.Handler(c)
	if err != nil {
		return NewErrorResponse(req.ID, err)
	}

	return NewResponse(req.ID, result)
}
