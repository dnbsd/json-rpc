package jsonrpc

import "encoding/json"

const version = "2.0"

var (
	ErrParseError     = NewError(-32700, "Parse error")
	ErrInvalidRequest = NewError(-32600, "Invalid Request")
	ErrMethodNotFound = NewError(-32601, "Method not found")
	ErrInvalidParams  = NewError(-32602, "Invalid params")
	ErrInternalError  = NewError(-32603, "Internal error")
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}

type Request struct {
	Version string `json:"jsonrpc"`
	ID      any    `json:"id,omitempty"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
}

func NewRequest(id any, method string, params any) Request {
	return Request{
		Version: version,
		ID:      id,
		Method:  method,
		Params:  params,
	}
}

func ParseRequest(b []byte) (Request, error) {
	var req Request
	err := json.Unmarshal(b, &req)
	if err != nil {
		return Request{}, ErrParseError
	}
	return req, nil
}

func (r Request) IsNotification() bool {
	return r.ID == nil
}

func (r Request) Validate() error {
	if r.Version != version {
		return ErrInvalidRequest
	}

	if r.Method == "" {
		return ErrInvalidRequest
	}

	switch r.Params.(type) {
	case map[string]any:
	case []any:
	default:
		return ErrInvalidParams
	}

	return nil
}
