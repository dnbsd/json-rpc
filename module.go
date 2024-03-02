package jsonrpc

import "context"

type Module interface {
	Exports() map[string]Handler
	Submodules() map[string]Module
}

type Handler func(context.Context, Params) (any, error)
