package metadata

import "fmt"

type TypeError struct {
	Key      any
	Expected string // 例: "int64"
	Actual   string // 例: "string"
}

func (e *TypeError) Error() string {
	return fmt.Sprintf("metadata: type mismatch for key %v: expected %s, got %s", e.Key, e.Expected, e.Actual)
}
