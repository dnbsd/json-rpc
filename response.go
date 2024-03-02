package jsonrpc

type Response struct {
	Version string `json:"jsonrpc"`
	ID      any    `json:"id,omitempty"`
	Result  any    `json:"result,omitempty"`
	Error   error  `json:"error,omitempty"`
}

func NewResponse(id, result any) Response {
	return Response{
		Version: version,
		ID:      id,
		Result:  result,
	}
}

func NewErrorResponse(id any, err error) Response {
	return Response{
		Version: version,
		ID:      id,
		Error:   err,
	}
}
