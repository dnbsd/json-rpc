package jsonrpc

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
