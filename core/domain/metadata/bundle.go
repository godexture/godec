package metadata

type Bundle struct {
	Tags map[string]string
	Raw  map[string]any
}

func NewBundle() *Bundle {
	return &Bundle{
		Tags: make(map[string]string),
		Raw:  make(map[string]any),
	}
}
