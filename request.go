package jsonrpc

const version = "2.0"

type Request struct {
	Version string `json:"jsonrpc"`
	ID      any    `json:"id,omitempty"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

func NewRequest(id uint64, method string, params Params) Request {
	return Request{
		Version: version,
		ID:      &id,
		Method:  method,
		Params:  params,
	}
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

	return nil
}
