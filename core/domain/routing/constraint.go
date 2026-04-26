package routing

import "github.com/godexture/core/domain/media"

type Constraint interface {
	MediaType() media.MediaType
	Matches(p media.Profile) bool
	Diagnose(p media.Profile) bool
}
