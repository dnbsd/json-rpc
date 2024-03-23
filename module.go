package jsonrpc

type Module struct {
	Name       string
	Methods    []Method
	Submodules []Module
}

type Method struct {
	Name    string
	Handler func(*Context) (any, error)
}
