package media

import (
	"math/bits"
	"slices"
)

type ChannelCategory uint8
type ChannelPosition uint64

const (
	OrderUnspecified ChannelCategory = iota
	OrderNative
	OrderCustom
	OrderAmbisonic
)

const (
	FrontLeft ChannelPosition = 1 << iota
	FrontRight
	FrontCenter
	LowFrequency
	BackLeft
	BackRight
	FrontLeftOfCenter
	FrontRightOfCenter
	BackCenter
	SideLeft
	SideRight
	TopCenter
	TopFrontLeft
	TopFrontCenter
	TopFrontRight
	TopBackLeft
	TopBackCenter
	TopBackRight
)

var (
	LayoutMono1        ChannelLayout = NewNativeLayout(FrontCenter)
	LayoutDualMono2    ChannelLayout = NewCustomLayout(FrontLeft | FrontRight)
	LayoutStereo2_0    ChannelLayout = NewNativeLayout(FrontLeft | FrontRight)
	LayoutStereo2_1    ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | LowFrequency)
	LayoutStereo3_0    ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter)
	LayoutStereo3_1    ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency)
	LayoutSurround3_0  ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | BackCenter)
	LayoutQuad4_0      ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | BackLeft | BackRight)
	LayoutSideQuad4_0  ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | SideLeft | SideRight)
	LayoutSurround4_0  ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | BackLeft | BackRight | FrontCenter)
	LayoutSurround4_1  ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | BackLeft | BackRight | FrontCenter | LowFrequency)
	LayoutFront5_0     ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | BackLeft | BackRight)
	LayoutFront5_1     ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | BackLeft | BackRight | LowFrequency)
	LayoutSide5_0      ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | SideLeft | SideRight | FrontCenter)
	LayoutSide5_1      ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | SideLeft | SideRight | FrontCenter | LowFrequency)
	LayoutAtmos5_1_4   ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight | TopFrontLeft | TopFrontCenter | TopFrontRight | TopBackLeft | TopBackCenter | TopBackRight)
	LayoutHexagonal6_0 ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | BackLeft | BackRight | BackCenter)
	LayoutHexagonal6_1 ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight | BackCenter)
	LayoutFront6_0     ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight)
	LayoutSide6_0      ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | SideLeft | SideRight | FrontCenter | LowFrequency)
	LayoutSide6_1      ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | SideLeft | SideRight | FrontCenter | LowFrequency | BackCenter)
	LayoutWide7_1      ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight | SideLeft | SideRight)
	LayoutSide7_0      ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | SideLeft | SideRight | FrontCenter | BackLeft | BackRight)
	LayoutSide7_1      ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | SideLeft | SideRight | FrontCenter | BackLeft | BackRight | LowFrequency)
	LayoutSurround7_1  ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight | SideLeft | SideRight | TopCenter)
	LayoutAtmos7_1_4   ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight | SideLeft | SideRight | TopFrontLeft | TopFrontCenter | TopFrontRight | TopBackLeft | TopBackCenter | TopBackRight)
	LayoutOctagonal8_0 ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight | SideLeft | SideRight | TopCenter)
	LayoutSurround9_0  ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight | SideLeft | SideRight | TopCenter | TopFrontLeft)
	LayoutSurround9_1  ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight | SideLeft | SideRight | TopCenter | TopFrontLeft | LowFrequency)
	LayoutSurround11_0 ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight | SideLeft | SideRight | TopCenter | TopFrontLeft | TopFrontRight)
	LayoutSurround11_1 ChannelLayout = NewNativeLayout(FrontLeft | FrontRight | FrontCenter | LowFrequency | BackLeft | BackRight | SideLeft | SideRight | TopCenter | TopFrontLeft | TopFrontRight | LowFrequency)
	LayoutAtmos11_1_4  ChannelLayout = NewCustomLayout(FrontLeft, FrontRight, FrontCenter, LowFrequency, BackLeft, BackRight, SideLeft, SideRight, TopFrontLeft, TopFrontCenter, TopFrontRight, TopBackLeft, TopBackCenter, TopBackRight)
	LayoutAtmos11_1_6  ChannelLayout = NewCustomLayout(FrontLeft, FrontRight, FrontCenter, LowFrequency, BackLeft, BackRight, SideLeft, SideRight, TopFrontLeft, TopFrontCenter, TopFrontRight, TopBackLeft, TopBackCenter, TopBackRight)
)

type ChannelLayout struct {
	order        ChannelCategory
	channelCount int
	mask         ChannelPosition
	custom       []ChannelPosition
}

func NewUnspecified(channels int) ChannelLayout {
	return ChannelLayout{
		order:        OrderUnspecified,
		channelCount: channels,
	}
}
func NewNativeLayout(mask ChannelPosition) ChannelLayout {
	return ChannelLayout{
		order:        OrderNative,
		channelCount: bits.OnesCount64(uint64(mask)),
		mask:         mask,
	}
}

func NewCustomLayout(channels ...ChannelPosition) ChannelLayout {
	var mask ChannelPosition

	for _, channel := range channels {
		mask |= channel
	}

	return ChannelLayout{
		order:        OrderCustom,
		channelCount: len(channels),
		mask:         mask,
		custom:       channels,
	}
}

func NewAmbisonicLayout(order uint8) ChannelLayout {
	channelCount := int((order + 1) * (order + 1))
	mask := ChannelPosition((1 << channelCount) - 1)
	return ChannelLayout{
		order:        OrderAmbisonic,
		channelCount: channelCount,
		mask:         mask,
	}
}

func (l ChannelLayout) IsUnspecified() bool {
	return l.order == OrderUnspecified
}

func (l ChannelLayout) IsSpatial() bool {
	return l.order == OrderNative || l.order == OrderCustom
}

func (l ChannelLayout) IsAmbisonic() bool {
	return l.order == OrderAmbisonic
}

func (l ChannelLayout) ChannelCount() int { return l.channelCount }

func (l ChannelLayout) Equals(other ChannelLayout) bool {
	return l.order == other.order &&
		l.channelCount == other.channelCount &&
		l.mask == other.mask &&
		slices.Equal(l.custom, other.custom)
}

func (l ChannelLayout) Contains(c ChannelPosition) bool {
	switch l.order {

	case OrderNative:
		return (l.mask & c) != 0

	case OrderCustom:
		return slices.Contains(l.custom, c)

	default:
		return false
	}
}

func (l ChannelLayout) Enumerate() []ChannelPosition {
	switch l.order {
	case OrderNative:
		var channels []ChannelPosition
		for i := 0; i < 64; i++ {
			if l.mask&(1<<i) != 0 {
				channels = append(channels, 1<<i)
			}
		}
		return channels

	case OrderCustom:
		return append([]ChannelPosition{}, l.custom...)

	default:
		return nil
	}
}

func (l ChannelLayout) Index(c ChannelPosition) int {
	switch l.order {

	case OrderNative:
		if (l.mask & c) == 0 {
			return -1
		}

		return bits.OnesCount64(uint64(l.mask & (c - 1)))

	case OrderCustom:
		return slices.Index(l.custom, c)

	default:
		return -1

	}
}

func (l ChannelLayout) String() string {
	return "custom(" + string(rune(l.ChannelCount()+'0')) + "ch)"
}
