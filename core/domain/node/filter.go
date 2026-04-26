package node

import (
	"context"

	"github.com/godexture/core/domain/media"
)

type Filter interface {
	Lifecycle

	Process(ctx context.Context, in <-chan media.Frame) (out <-chan media.Frame, errs <-chan error)
}
