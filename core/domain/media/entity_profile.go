package media

type MediaType string

const (
	MediaTypeUnknown    MediaType = ""
	MediaTypeVideo      MediaType = "video"
	MediaTypeAudio      MediaType = "audio"
	MediaTypeSubtitle   MediaType = "subtitle"
	MediaTypeData       MediaType = "data"
	MediaTypeAttachment MediaType = "attachment"
)

type Profile interface {
	Type() MediaType
	Equals(Profile) bool
}
