package media

type AudioProfile struct {
	SampleRate    int
	ChannelLayout ChannelLayout
	Format        SampleFormat
}

func (p AudioProfile) Type() MediaType {
	return MediaTypeAudio
}

func (p AudioProfile) Equals(other Profile) bool {
	audioProfile, ok := other.(*AudioProfile)

	return ok && p.SampleRate == audioProfile.SampleRate &&
		p.ChannelLayout.Equals(audioProfile.ChannelLayout) &&
		p.Format == audioProfile.Format
}
