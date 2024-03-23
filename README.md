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

builder := jsonrpc.Builder{
    Modules: []jsonrpc.Module{
        {
            Name: "math",
            Methods: []jsonrpc.Method{
                {
                    Name: "version",
                    Handler: func(c *jsonrpc.Context) (any, error) {
                        return c.Result("v1.0")
                    },
                },
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
            },
        },
    },
}

service, err := builder.Build()
if err != nil {
    panic(err)
}

req := jsonrpc.NewRequest(1, "math.add", jsonrpc.ParamsArray{10, 20})
err := req.Validate()
if err != nil {
    panic(err)
}

resp := service.Call(context.Background(), req)
if resp.Error != nil {
    panic(err)
}

println("sum", resp.Result.(int))
```

## License

MIT