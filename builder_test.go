package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilderBuild(t *testing.T) {
	builder := Builder{
		Modules: []Module{
			{
				Name: "std",
				Submodules: []Module{
					{
						Name: "math",
						Methods: []Method{
							{
								Name: "add",
								Handler: func(c *Context) (any, error) {
									params, err := c.ParamsArray()
									if err != nil {
										return c.Error(err)
									}

									var sum int
									for i := range params {
										n, err := params.Number(i)
										if err != nil {
											return c.Error(err)
										}
										sum += n.Int()
									}

									return c.Result(sum)
								},
							},
							{
								Name: "divide",
								Handler: func(c *Context) (any, error) {
									return nil, nil
								},
							},
						},
					},
					{
						Name: "string",
						Methods: []Method{
							{
								Name: "join",
								Handler: func(c *Context) (any, error) {
									return nil, nil
								},
							},
							{
								Name: "split",
								Handler: func(c *Context) (any, error) {
									return nil, nil
								},
							},
						},
					},
				},
			},
		},
	}
	_, err := builder.Build()
	assert.NoError(t, err)
}

func TestBuilderBuildMissingModuleName(t *testing.T) {
	builder := &Builder{
		Modules: []Module{
			{
				Name: "",
			},
		},
	}
	_, err := builder.Build()
	assert.ErrorContains(t, err, "module has no name")
}

func TestBuilderBuildMissingSubmoduleName(t *testing.T) {
	builder := &Builder{
		Modules: []Module{
			{
				Name: "std",
				Submodules: []Module{
					{
						Name: "",
					},
				},
			},
		},
	}
	_, err := builder.Build()
	assert.ErrorContains(t, err, "std: module has no name")
}

func TestBuilderBuildMissingModuleMethodName(t *testing.T) {
	builder := &Builder{
		Modules: []Module{
			{
				Name: "std",
				Methods: []Method{
					{
						Name: "",
						Handler: func(c *Context) (any, error) {
							return nil, nil
						},
					},
				},
			},
		},
	}
	_, err := builder.Build()
	assert.ErrorContains(t, err, "std: method has no name")
}

func TestBuilderBuildMissingSubmoduleMethodName(t *testing.T) {
	builder := &Builder{
		Modules: []Module{
			{
				Name: "std",
				Submodules: []Module{
					{
						Name: "math",
						Methods: []Method{
							{
								Name: "",
								Handler: func(c *Context) (any, error) {
									return nil, nil
								},
							},
						},
					},
				},
			},
		},
	}
	_, err := builder.Build()
	assert.ErrorContains(t, err, "std.math: method has no name")
}
