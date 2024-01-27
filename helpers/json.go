package helpers

import "encoding/json"

func UnsafeMarshal(v any) string {
	res, _ := json.Marshal(v)

	return string(res)
}

func UnsafeUnmarshal[T any](s string) T {
	res := new(T)
	json.Unmarshal([]byte(s), res)

	return *res
}
