package main

import (
	_ "embed"
)

var (
	//go:embed shaders/default.frag
	DefaultFrag []byte

	//go:embed shaders/default.vert
	DefaultVert []byte
)

// ----------------------------------------------------------------------------
func getShaderProgram(srcBytes []byte) string {
	return string(srcBytes)
}
