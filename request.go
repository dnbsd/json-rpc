package jsonrpc

import (
	"errors"
)

const version = "2.0"

var (
	ErrUnsupportedVersion = errors.New("unsupported version")
	ErrEmptyMethod        = errors.New("empty method")
)

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
		return ErrUnsupportedVersion
	}

	if r.Method == "" {
		return ErrEmptyMethod
	}

	return nil
}
