package media

import "github.com/godexture/core/domain/time"

type Stream struct {
	Index     int
	Type      MediaType
	CodecName string
	Timebase  time.Rational
}
