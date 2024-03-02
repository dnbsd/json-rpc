package jsonrpc

import "context"

type Module interface {
	Exports() map[string]Handler
	// TODO
	//Submodules() map[string]Module
}

type Handler func(context.Context, Params) (any, error)
