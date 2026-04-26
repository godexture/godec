package node

import (
	"github.com/godexture/core/domain/media"
	"github.com/godexture/core/domain/metadata"
	"github.com/godexture/core/domain/time"
)

type Demuxer interface {
	Lifecycle

	Streams() []media.Stream
	ReadPacket() (media.Packet, error)
	Metadata() *metadata.Bundle
}

type Muxer interface {
	Lifecycle

	AddStream(codecName string, tb time.Rational) (streamIndex int, err error)
	WritePacket(streamIndex int, pkt media.Packet) error
	SetMetadata(meta *metadata.Bundle) error
}
