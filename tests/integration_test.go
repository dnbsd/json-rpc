package tests

import (
	"context"
	"github.com/dnbsd/jsonrpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceModuleMethodCall(t *testing.T) {
	req, err := jsonrpc.ParseRequest([]byte(`{"jsonrpc": "2.0", "method": "std.version", "id": 1}`))
	assert.NoError(t, err)

	service := buildService(t)
	resp := service.Call(context.Background(), req)
	assert.Equal(t, float64(1), resp.ID)
	assert.Equal(t, float64(1), resp.Result)
	assert.Nil(t, resp.Error)
}

func TestServiceModuleMethodCallParamsObject(t *testing.T) {
	req, err := jsonrpc.ParseRequest([]byte(`{"jsonrpc": "2.0", "method": "std.echo", "params": { "message": "hello" }, "id": 1}`))
	assert.NoError(t, err)

	service := buildService(t)
	resp := service.Call(context.Background(), req)
	assert.Equal(t, float64(1), resp.ID)
	assert.Equal(t, map[string]any{"message": "hello"}, resp.Result)
	assert.Nil(t, resp.Error)
}

func TestServiceModuleMethodCallUndefined(t *testing.T) {
	req, err := jsonrpc.ParseRequest([]byte(`{"jsonrpc": "2.0", "method": "std.xxx", "id": 1}`))
	assert.NoError(t, err)

	service := buildService(t)
	resp := service.Call(context.Background(), req)
	assert.Equal(t, float64(1), resp.ID)
	assert.Equal(t, -32601, resp.Error.Code)
}

func TestServiceSubmoduleMethodCall(t *testing.T) {
	req, err := jsonrpc.ParseRequest([]byte(`{"jsonrpc": "2.0", "method": "std.math.add", "params": [42, 23], "id": 1}`))
	assert.NoError(t, err)

	service := buildService(t)
	resp := service.Call(context.Background(), req)
	assert.Equal(t, float64(1), resp.ID)
	assert.Equal(t, 65, resp.Result)
	assert.Nil(t, resp.Error)
}

func TestServiceSubmoduleMethodCallUndefined(t *testing.T) {
	req, err := jsonrpc.ParseRequest([]byte(`{"jsonrpc": "2.0", "method": "std.math.subtract", "params": [42, 23], "id": 1}`))
	assert.NoError(t, err)

	service := buildService(t)
	resp := service.Call(context.Background(), req)
	assert.Equal(t, float64(1), resp.ID)
	assert.Equal(t, -32601, resp.Error.Code)
}

func buildService(t *testing.T) *jsonrpc.Service {
	builder := jsonrpc.Builder{
		Modules: []jsonrpc.Module{
			{
				Name: "std",
				Methods: []jsonrpc.Method{
					{
						Name: "version",
						Handler: func(c *jsonrpc.Context) (any, error) {
							return 1.0, nil
						},
					},
					{
						Name: "echo",
						Handler: func(c *jsonrpc.Context) (any, error) {
							params, err := c.ParamsObject()
							if err != nil {
								return c.Error(err)
							}

							msg, err := params.String("message")
							if err != nil {
								return c.Error(err)
							}

							return c.Result(map[string]any{
								"message": msg,
							})
						},
					},
				},
				Submodules: []jsonrpc.Module{
					{
						Name: "math",
						Methods: []jsonrpc.Method{
							{
								Name: "add",
								Handler: func(c *jsonrpc.Context) (any, error) {
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
								Handler: func(c *jsonrpc.Context) (any, error) {
									return nil, nil
								},
							},
						},
					},
					{
						Name: "string",
						Methods: []jsonrpc.Method{
							{
								Name: "join",
								Handler: func(c *jsonrpc.Context) (any, error) {
									return nil, nil
								},
							},
							{
								Name: "split",
								Handler: func(c *jsonrpc.Context) (any, error) {
									return nil, nil
								},
							},
						},
					},
				},
			},
		},
	}
	s, err := builder.Build()
	assert.NoError(t, err)
	return s
}
