package assets

import (
	"embed"
	"io"

	resource "github.com/quasilyte/ebitengine-resource"
)

//go:embed all:_data
var gameAssets embed.FS

func MakeOpenAssetFunc() func(path string) io.ReadCloser {
	return func(path string) io.ReadCloser {
		f, err := gameAssets.Open("_data/" + path)
		if err != nil {
			panic(err)
		}
		return f
	}
}

func RegisterResources(loader *resource.Loader) {
	registerImageResources(loader)
	registerAudioResources(loader)
	registerShaderResources(loader)
}
