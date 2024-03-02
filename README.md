# json-rpc

[![GoDoc][GoDoc-Image]][GoDoc-Url]

[GoDoc-Image]: https://img.shields.io/badge/GoDoc-reference-007d9c
[GoDoc-Url]: https://pkg.go.dev/github.com/dnbsd/jsonrpc

## Installation

```
go get github.com/dnbsd/jsonrpc
```

## Usage

```go
import "github.com/dnbsd/jsonrpc"

type EchoModule struct{}

func (m *EchoModule) Echo(ctx context.Context, params jsonrpc.Params) (any, error) {
    message := params.String("message")
    if message == "" {
        return nil, errors.New("empty message")
    }
    return message, nil
}

func (m *EchoModule) Submodules() map[string]jsonrpc.Module {
    return nil
}

func (m *EchoModule) Exports() map[string]jsonrpc.Method {
    return map[string]jsonrpc.Method{
        "echo": m.Echo,
    }
}

rpc := jsonrpc.New()
rpc.Register("echo", &EchoModule{})
req := jsonrpc.NewRequest(1, "echo.echo", jsonrpc.Params{
    "message": "hello world",
})
resp := rpc.Call(context.Background(), req)
```

## License

MIT