package main

import "github.com/gowebapi/webapi/graphics/webgl"

// ----------------------------------------------------------------------------
func setupShaders(gl *webgl.RenderingContext) *webgl.Program {
	vertexShaderCode := getShaderProgram(DefaultVert)
	vertexShader := gl.CreateShader(webgl.VERTEX_SHADER)
	gl.ShaderSource(vertexShader, vertexShaderCode)
	gl.CompileShader(vertexShader)

	fragmentShaderCode := getShaderProgram(DefaultFrag)
	fragmentShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fragmentShader, fragmentShaderCode)
	gl.CompileShader(fragmentShader)

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)

	gl.UseProgram(prog)

	modelviewLoc := gl.GetUniformLocation(prog, "modelview")

	rotationMatrix := rotateOnZAxis(angle)
	gl.UniformMatrix4fv(modelviewLoc, false, rotationMatrix)

	return prog
}
