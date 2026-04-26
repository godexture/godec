package converter

import (
	"github.com/godexture/core/domain/media"
	"github.com/godexture/core/domain/node"
)

type errorAwareDecoder struct {
	base    node.Decoder
	handler media.ErrorHandler
}

func NewErrorAwareDecoder(d node.Decoder, h media.ErrorHandler) node.Decoder {
	return &errorAwareDecoder{base: d, handler: h}
}

func (d *errorAwareDecoder) Decode(pkt media.Packet) (media.Frame, error) {
	frame, err := d.base.Decode(pkt)
	if err != nil {
		wrapped := &media.Error{
			Stage:       media.StageDecoder,
			StreamIndex: pkt.StreamIndex(),
			PTS:         pkt.Pts(),
			Err:         err,
		}

		if d.handler != nil && d.handler.Handle(wrapped) == media.ActionIgnore {
			return nil, nil
		}
		return nil, wrapped
	}
	return frame, nil
}

func (d *errorAwareDecoder) Init() error {
	return d.base.Init()
}

func (d *errorAwareDecoder) Flush() error {
	return d.base.Flush()
}

func (d *errorAwareDecoder) Close() error {
	return d.base.Close()
}
