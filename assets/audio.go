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
	audioResources := map[resource.AudioID]resource.AudioInfo{
		AudioBowShot1: {Path: "audio/bow_shot1.wav"},
		AudioBowShot2: {Path: "audio/bow_shot2.wav"},
		AudioBowShot3: {Path: "audio/bow_shot3.wav", Volume: -0.6},

		AudioBluntAttack1: {Path: "audio/blunt_attack1.wav"},
		AudioBluntAttack2: {Path: "audio/blunt_attack2.wav"},

		AudioSwordAttack1: {Path: "audio/sword_attack1.wav"},
		AudioSwordAttack2: {Path: "audio/sword_attack2.wav"},

		AudioDeath1: {Path: "audio/death1.wav"},
	}

	for id, res := range audioResources {
		loader.AudioRegistry.Set(id, res)
		loader.LoadAudio(id)
	}
}

func NumSamples(a resource.AudioID) int {
	switch a {
	case AudioBowShot1:
		return 3
	case AudioBluntAttack1:
		return 2
	case AudioSwordAttack1:
		return 2
	default:
		return 1
	}
}

const (
	AudioNone resource.AudioID = iota

	AudioBowShot1
	AudioBowShot2
	AudioBowShot3
	AudioBluntAttack1
	AudioBluntAttack2
	AudioSwordAttack1
	AudioSwordAttack2
	AudioDeath1
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
