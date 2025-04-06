//kage:unit pixels

//go:build ignore

package main

var Time float

func Fragment(_ vec4, pos vec2, _ vec4) vec4 {
	alpha := imageSrc0UnsafeAt(pos).a
	noise := 1.1 * imageSrc1UnsafeAt(pos).r

	const threshold = 0.1

	if noise < threshold {
		return vec4(0)
	}

	meltAmount := (noise - threshold) * 20.0
	meltedCoord := pos + vec2(sin(17*(meltAmount*(1.0+(Time*0.55)))), 16*(meltAmount*-(Time*0.9)))

	meltedColor := imageSrc0At(meltedCoord)

	m := alpha * (1.0 - (Time * 1.1))
	meltedColor.r *= (0.9 + (Time * 0.5))
	return meltedColor * m
}
