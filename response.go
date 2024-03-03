package jsonrpc

type Response struct {
	Version string `json:"jsonrpc"`
	ID      any    `json:"id,omitempty"`
	Result  any    `json:"result,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

func NewResponse(id, result any) Response {
	return Response{
		Version: version,
		ID:      id,
		Result:  result,
	}
}

func NewErrorResponse(id any, err error) Response {
	respError, ok := err.(*Error)
	if !ok {
		respError = NewError(1, err.Error())
	}
	return Response{
		Version: version,
		ID:      id,
		Error:   respError,
	}
}
