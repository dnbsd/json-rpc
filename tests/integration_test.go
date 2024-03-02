package tests

import (
	"context"
	"errors"
	"github.com/dnbsd/jsonrpc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type EchoSubmodule struct{}

func (m *EchoSubmodule) UpperCase(ctx context.Context, params jsonrpc.Params) (any, error) {
	message := params.String("message")
	if message == "" {
		return nil, errors.New("empty message")
	}
	return strings.ToUpper(message), nil
}

func (m *EchoSubmodule) Exports() map[string]jsonrpc.Method {
	return map[string]jsonrpc.Method{
		"upper_case": m.UpperCase,
	}
}

func (m *EchoSubmodule) Submodules() map[string]jsonrpc.Module {
	return nil
}

type EchoModule struct{}

func (m *EchoModule) Echo(ctx context.Context, params jsonrpc.Params) (any, error) {
	message := params.String("message")
	if message == "" {
		return nil, errors.New("empty message")
	}
	return message, nil
}

func (m *EchoModule) Submodules() map[string]jsonrpc.Module {
	return map[string]jsonrpc.Module{
		"sub": &EchoSubmodule{},
	}
}

func (m *EchoModule) Exports() map[string]jsonrpc.Method {
	return map[string]jsonrpc.Method{
		"echo": m.Echo,
	}
}

func TestMethodCall(t *testing.T) {
	rpc := jsonrpc.New()
	rpc.Register("echo", &EchoModule{})
	req := jsonrpc.NewRequest(1, "echo.echo", jsonrpc.Params{
		"message": "hello world",
	})
	resp := rpc.Call(context.Background(), req)
	assert.NoError(t, resp.Error)
	assert.Equal(t, resp.Result, "hello world")
	assert.Equal(t, req.ID, resp.ID)
}

func TestMethodCallUndefined(t *testing.T) {
	rpc := jsonrpc.New()
	rpc.Register("echo", &EchoModule{})
	req := jsonrpc.NewRequest(1, "echo.echos", jsonrpc.Params{
		"message": "hello world",
	})
	resp := rpc.Call(context.Background(), req)
	assert.ErrorIs(t, jsonrpc.ErrMethodNotFound, resp.Error)
	assert.Equal(t, resp.Result, nil)
	assert.Equal(t, req.ID, resp.ID)
}

func TestSubmoduleMethodCall(t *testing.T) {
	rpc := jsonrpc.New()
	rpc.Register("echo", &EchoModule{})
	req := jsonrpc.NewRequest(1, "echo.sub.upper_case", jsonrpc.Params{
		"message": "hello world",
	})
	resp := rpc.Call(context.Background(), req)
	assert.NoError(t, resp.Error)
	assert.Equal(t, resp.Result, "HELLO WORLD")
	assert.Equal(t, req.ID, resp.ID)
}
