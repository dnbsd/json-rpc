package jsonrpc

import (
	"errors"
	"strings"
)

type Builder struct {
	Modules []Module
	methods map[string]Method
}

func (b *Builder) buildModule(c BuilderContext, module Module) error {
	if module.Name == "" {
		return c.NewError("module has no name")
	}

	c = c.Add(module.Name)

	for _, method := range module.Methods {
		if method.Name == "" {
			return c.NewError("method has no name")
		}
		fqn := c.FQN(method.Name)
		b.methods[fqn] = method
	}

	for _, submodule := range module.Submodules {
		err := b.buildModule(c, submodule)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Builder) Build() (*Service, error) {
	b.methods = make(map[string]Method)
	for _, module := range b.Modules {
		err := b.buildModule(BuilderContext{}, module)
		if err != nil {
			return nil, err
		}
	}
	return &Service{
		methods: b.methods,
	}, nil
}

type BuilderContext struct {
	path []string
}

func (c BuilderContext) Add(name string) BuilderContext {
	c.path = append(c.path, name)
	return c
}

func (c BuilderContext) FQN(method string) string {
	var fqn []string
	fqn = append(fqn, c.path...)
	fqn = append(fqn, method)
	return strings.Join(fqn, ".")
}

func (c BuilderContext) NewError(s string) error {
	if len(c.path) == 0 {
		return errors.New(s)
	}
	path := strings.Join(c.path, ".")
	return errors.New(path + ": " + s)
}
