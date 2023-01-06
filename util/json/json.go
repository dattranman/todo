package json

import "encoding/json"

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmashal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
