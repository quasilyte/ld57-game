//go:build ignore

//kage:unit pixels

package main

var Angle float

func Fragment(_ vec4, pos vec2, clr2 vec4) vec4 {
	clr := imageSrc0UnsafeAt(pos)

	// -1.5
	rotated := hueRotate(clr.rgb, Angle) * 1.05
	return vec4(rotated, clr.a) * clr2
}

func hueRotate(color vec3, angle float) vec3 {
	rad := angle
	cosA := cos(rad)
	sinA := sin(rad)

	hueRotationMatrix := mat3(
		vec3(0.299, 0.587, 0.114),
		vec3(0.299, 0.587, 0.114),
		vec3(0.299, 0.587, 0.114),
	) +
		mat3(
			vec3(0.701, -0.587, -0.114),
			vec3(-0.299, 0.413, -0.114),
			vec3(-0.3, -0.588, 0.886),
		)*cosA +
		mat3(
			vec3(0.168, 0.330, -0.497),
			vec3(-0.328, 0.035, 0.292),
			vec3(1.25, -1.05, -0.203),
		)*sinA

	return clamp(color*hueRotationMatrix, 0.0, 1.0)
}
