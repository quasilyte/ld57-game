package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"

	_ "image/png"
)

const (
	SoundGroupEffect uint = iota
	SoundGroupMusic
)

func registerAudioResources(loader *resource.Loader) {
	audioResources := map[resource.AudioID]resource.AudioInfo{}

	for id, res := range audioResources {
		loader.AudioRegistry.Set(id, res)
		loader.LoadAudio(id)
	}
}

func NumSamples(a resource.AudioID) int {
	switch a {
	default:
		return 1
	}
}

const (
	AudioNone resource.AudioID = iota
)

func VolumeMultiplier(level int) float64 {
	switch level {
	case 1:
		return 0.01
	case 2:
		return 0.15
	case 3:
		return 0.45
	case 4:
		return 0.8
	case 5:
		return 1.0
	default:
		return 0
	}
}
