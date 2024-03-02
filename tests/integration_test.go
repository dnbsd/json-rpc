package tests

import (
	"context"
	"errors"
	"github.com/dnbsd/jsonrpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

type EchoModule struct{}

func (m *EchoModule) Echo(ctx context.Context, params jsonrpc.Params) (any, error) {
	message := params.String("message")
	if message == "" {
		return nil, errors.New("empty message")
	}
	return message, nil
}

func (m *EchoModule) Exports() map[string]jsonrpc.Handler {
	return map[string]jsonrpc.Handler{
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
