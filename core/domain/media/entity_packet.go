package media

import "github.com/godexture/core/domain/time"

type Packet interface {
	Retainer
	MediaType() MediaType
	StreamIndex() int
	Data() []byte
	Pts() Pts
	Dts() Dts
	Timebase() time.Rational
}
