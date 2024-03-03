package jsonrpc

import (
	"context"
	"strings"
)

type RPC struct {
	methods map[string]Method
}

func New() *RPC {
	return &RPC{
		methods: make(map[string]Method),
	}
}

// Register registers a module and its methods. If a module has submodules, all their methods are registered
// recursively. Registered methods are callable as `module.method`, or `module.submodule.method`.
func (s *RPC) Register(name string, module Module) {
	for methodName, method := range module.Exports() {
		fqn := strings.Join([]string{name, methodName}, ".")
		s.methods[fqn] = method
	}

	for submoduleName, submodule := range module.Submodules() {
		fqn := strings.Join([]string{name, submoduleName}, ".")
		s.Register(fqn, submodule)
	}
}

// Call calls a registered method using the provided request data. The method does not check if a request is a notification
// and always returns a response. It's callers responsibility to validate request data. If an unregistered method id called,
// the method returns ErrMethodNotFound. Otherwise, response data is populated by values returned from a method.
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
