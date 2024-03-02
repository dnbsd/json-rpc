package jsonrpc

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrMethodNotFound = errors.New("method not found")
)

type RPC struct {
	methods map[string]Handler
}

func New() *RPC {
	return &RPC{
		methods: make(map[string]Handler),
	}
}

func (s *RPC) Register(name string, module Module) {
	for methodName, handler := range module.Exports() {
		fqn := strings.Join([]string{name, methodName}, ".")
		s.methods[fqn] = handler
	}
}

func (s *RPC) Call(ctx context.Context, req Request) Response {
	method, ok := s.methods[req.Method]
	if !ok {
		return NewErrorResponse(req.ID, ErrMethodNotFound)
	}

	result, err := method(ctx, req.Params)
	if err != nil {
		return NewErrorResponse(req.ID, err)
	}

	return NewResponse(req.ID, result)
}
