package media

type Frame interface {
	Retainer
	MediaType() MediaType
	Pts() Pts
	Profile() Profile
}
