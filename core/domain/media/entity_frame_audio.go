// core/domain/media/frame_audio.go
package media

import "unsafe"

type SampleType interface {
	~uint8 | ~int16 | ~int32 | ~float32 | ~float64
}

type AudioFrame struct {
	Format     SampleFormat
	Layout     ChannelLayout
	SampleRate int
	Samples    int

	planes [][]byte
}

func Plane[T SampleType](f *AudioFrame, planeIndex int) []T {
	if planeIndex >= len(f.planes) {
		return nil
	}
	bytes := f.planes[planeIndex]
	if len(bytes) == 0 {
		return nil
	}

	ptr := (*T)(unsafe.Pointer(&bytes[0]))

	length := len(bytes) / int(unsafe.Sizeof(*ptr))

	return unsafe.Slice(ptr, length)
}
