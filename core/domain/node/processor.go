package node

import "github.com/godexture/core/domain/media"

type Decoder interface {
	Lifecycle

	Decode(media.Packet) (media.Frame, error)
}

type Encoder interface {
	Lifecycle

	Encode(media.Frame) (media.Packet, error)
}

type AudioDecoder interface {
	Decoder
	OutputProfile() media.AudioProfile
}

type AudioEncoder interface {
	Encoder
	SupportedProfiles() []media.AudioProfile
}
