package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"

	_ "image/png"
)

func registerShaderResources(loader *resource.Loader) {
	resources := map[resource.ShaderID]resource.ShaderInfo{
		ShaderMelt:      {Path: "shader/melt.go"},
		ShaderHueRotate: {Path: "shader/huerotate.go"},
	}

	for id, info := range resources {
		loader.ShaderRegistry.Set(id, info)
		loader.LoadShader(id)
	}
}

const (
	ShaderNone resource.ShaderID = iota

	ShaderMelt
	ShaderHueRotate
)
