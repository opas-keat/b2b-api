package shareutils

import "encoding/json"

func FromJson[T any](bytes []byte) (*T, error) {
	out := new(T)
	if err := json.Unmarshal(bytes, out); err != nil {
		return nil, err
	}
	return out, nil
}

func FromJsonNotErr[T any](bytes []byte) *T {
	out := new(T)
	if err := json.Unmarshal(bytes, out); err != nil {
		return nil
	}
	return out
}

func ToJson(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func ToJsonNotErr(v interface{}) []byte {
	out, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return out
}
