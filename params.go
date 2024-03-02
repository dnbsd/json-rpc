package jsonrpc

type Params map[string]any

func (p Params) Exists(key string) bool {
	_, ok := p[key]
	return ok
}

func (p Params) String(key string) string {
	if !p.Exists(key) {
		return ""
	}
	v, ok := p[key].(string)
	if !ok {
		return ""
	}
	return v
}

func (p Params) Float64(key string) float64 {
	if !p.Exists(key) {
		return 0
	}
	v, ok := p[key].(float64)
	if !ok {
		return 0
	}
	return v
}

func (p Params) Value(key string) any {
	if !p.Exists(key) {
		return nil
	}
	return p[key]
}
