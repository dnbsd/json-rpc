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

// Register registers a module and its handlers. If a module has submodules, all their handlers are registered
// recursively. Registered handlers are callable as `module.handler`, or `module.submodule.handler`.
func (s *RPC) Register(name string, module Module) {
	for methodName, handler := range module.Exports() {
		fqn := strings.Join([]string{name, methodName}, ".")
		s.methods[fqn] = handler
	}

	for submoduleName, submodule := range module.Submodules() {
		fqn := strings.Join([]string{name, submoduleName}, ".")
		s.Register(fqn, submodule)
	}
}

// Call calls a registered handler using the provided request data. The method does not check if a request is a notification
// and always returns a response. It's callers responsibility to validate request data. If an unregistered method id called,
// the method returns ErrMethodNotFound. Otherwise, response data is populated by values returned from a handler.
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
