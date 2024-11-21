package main

import (
	"math"

	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
)

// ----------------------------------------------------------------------------
// Rotates the matrix around the Z axis by the given angle
//
//	 Go code for:
//		mat4 rotationZ(float angle) {
//			return mat4(
//				cos(angle), -sin(angle), 0,0,
//				sin(angle), cos(angle), 0, 0,
//				0,0,1,0,
//				0,0,0,1);
//		}
func rotateOnZAxis(angle float32) *webgl.Union {
	cos := float32(math.Cos(float64(angle)))
	sin := float32(math.Sin(float64(angle)))

	return webgl.UnionFromJS(jsconv.Float32ToJs([]float32{
		cos, -sin, 0, 0,
		sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}))
}
