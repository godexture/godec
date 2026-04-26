package routing

import (
	"fmt"
	"slices"

	"github.com/godexture/core/domain/media"
)

type rejectReason uint8

const (
	reasonNone rejectReason = iota
	reasonTypeMismatch
	reasonInvalidProfile
	reasonSampleRate
	reasonChannel
	reasonLayout
	reasonFormat
)

type AudioConstraint struct {
	SampleRates []int
	Channels    []int
	Layouts     []media.ChannelLayout
	Formats     []media.SampleFormat
}

func (c *AudioConstraint) check(p media.Profile) rejectReason {
	if p.Type() != media.MediaTypeAudio {
		return reasonTypeMismatch
	}

	profile, ok := p.(*media.AudioProfile)
	if !ok {
		return reasonInvalidProfile
	}

	if !slices.Contains(c.SampleRates, profile.SampleRate) {
		return reasonSampleRate
	}

	if !slices.Contains(c.Channels, profile.ChannelLayout.ChannelCount()) {
		return reasonChannel
	}

	if !slices.Contains(c.Formats, profile.Format) {
		return reasonFormat
	}

	if !slices.ContainsFunc(c.Layouts, func(l media.ChannelLayout) bool {
		return l.Equals(profile.ChannelLayout)
	}) {
		return reasonLayout
	}

	return reasonNone
}

func (c *AudioConstraint) Match(p media.Profile) bool {
	return c.check(p) == reasonNone
}

func (c *AudioConstraint) Diagnose(p media.Profile) error {
	code := c.check(p)

	switch code {
	case reasonNone:
		return nil

	case reasonTypeMismatch:
		return fmt.Errorf("type mismatch: expected audio, got %s", p.Type())

	case reasonInvalidProfile:
		return fmt.Errorf("internal error: profile is not AudioProfile")

	case reasonSampleRate:
		ap := p.(*media.AudioProfile)
		return fmt.Errorf("unsupported sample rate: %d Hz (allowed: %v)",
			ap.SampleRate, c.SampleRates)

	case reasonChannel:
		ap := p.(*media.AudioProfile)
		return fmt.Errorf("unsupported channel: %d ch. (allowed: %v)",
			ap.ChannelLayout.ChannelCount(), c.Channels)

	case reasonLayout:
		ap := p.(*media.AudioProfile)
		return fmt.Errorf("unsupported channel layout: %s (allowed: %v)",
			ap.ChannelLayout.String(), c.Layouts)

	default:
		return fmt.Errorf("unknown constraint violation")
	}
}
