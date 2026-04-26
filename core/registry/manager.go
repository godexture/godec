// ============================================================================
// File: registry/format.go
// 役割: プラグインの自己登録機構と依存性解決（DI）
// ============================================================================
package registry

import (
	"errors"
	"io"

	"github.com/godexture/core/domain/node"
)

type DemuxerFactory func(r io.ReadSeeker) (node.Demuxer, error)
type ProbeFunc func(r io.ReadSeeker) bool

var (
	demuxers = make(map[string]DemuxerFactory)
	probes   = []struct {
		name  string
		probe ProbeFunc
	}{}
)

func RegisterDemuxer(name string, factory DemuxerFactory, probe ProbeFunc) {
	demuxers[name] = factory
	if probe != nil {
		probes = append(probes, struct {
			name  string
			probe ProbeFunc
		}{name, probe})
	}
}

func DetectFormat(r io.ReadSeeker) (string, error) {
	for _, p := range probes {
		if p.probe(r) {
			r.Seek(0, io.SeekStart)
			return p.name, nil
		}
		r.Seek(0, io.SeekStart)
	}
	return "", errors.New("unknown format")
}

func GetDemuxer(name string, r io.ReadSeeker) (node.Demuxer, error) {
	factory, ok := demuxers[name]
	if !ok {
		return nil, errors.New("unsupported format")
	}
	return factory(r)
}
