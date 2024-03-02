package jsonrpc

import "context"

type Module interface {
	Exports() map[string]Method
	Submodules() map[string]Module
}

type Method func(context.Context, Params) (any, error)
