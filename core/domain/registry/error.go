// core/domain/registry/error.go
package registry

import "fmt"

type ComponentType string

const (
	Decoder ComponentType = "decoder"
	Encoder ComponentType = "encoder"
	Demuxer ComponentType = "demuxer"
	Muxer   ComponentType = "muxer"
	Filter  ComponentType = "filter"
)

type Error struct {
	Component ComponentType
	Name      string
	Available []string
}

func (e *Error) Error() string {
	return fmt.Sprintf("requested %s '%s' is not registered", e.Component, e.Name)
}
