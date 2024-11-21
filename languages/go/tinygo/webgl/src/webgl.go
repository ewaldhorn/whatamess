package main

import (
	"fmt"
	"time"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/core/js"
	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
)

// ----------------------------------------------------------------------------
var angle float32

// ----------------------------------------------------------------------------
func drawFrame(this js.Value, p []js.Value) interface{} {
	startTime := time.Now()

	angle += 0.01

	clearScreen()
	setupViewport()

	rotationMatrix := rotateOnZAxis(angle)
	coord := glContext.GetAttribLocation(shaderProgram, "coordinates")

	glContext.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
	glContext.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)

	glContext.VertexAttribPointer(uint(coord), 3, webgl.FLOAT, false, 0, 0)
	glContext.EnableVertexAttribArray(uint(coord))

	modelviewLoc := glContext.GetUniformLocation(shaderProgram, "modelview")
	glContext.UniformMatrix4fv(modelviewLoc, false, rotationMatrix)

	glContext.DrawElements(webgl.TRIANGLES, indicesCount, webgl.UNSIGNED_SHORT, 0)

	elapsedSeconds := time.Now().Sub(startTime).Seconds()
	fps := float64(1.0 / elapsedSeconds)

	fpsDisplay := fmt.Sprintf("FPS: %08.2f", fps)
	doc := webapi.GetWindow().Document()
	fpsElem := doc.GetElementById("fps")
	fpsElem.SetTextContent(&fpsDisplay)

	js.Global().Call("requestAnimationFrame", js.FuncOf(drawFrame))

	return nil
}

// ----------------------------------------------------------------------------
func createBuffers(gl *webgl.RenderingContext) (*webgl.Buffer, *webgl.Buffer, int) {
	var verticesNative = []float32{-0.5, 0.5, 0, -0.5, -0.5, 0, 0.5, -0.5, 0}
	var vertices = jsconv.Float32ToJs(verticesNative)
	vBuffer := gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, vBuffer)
	gl.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(vertices), webgl.STATIC_DRAW)
	gl.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{}) // unbind

	var indicesNative = []uint32{2, 1, 0}
	var indices = jsconv.UInt32ToJs(indicesNative)

	iBuffer := gl.CreateBuffer()
	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, iBuffer)
	gl.BufferData2(webgl.ELEMENT_ARRAY_BUFFER, webgl.UnionFromJS(indices), webgl.STATIC_DRAW)
	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, &webgl.Buffer{}) // unbind

	return vBuffer, iBuffer, len(indicesNative)
}

// ----------------------------------------------------------------------------
func clearScreen() {
	glContext.ClearColor(0.05, 0.05, 0.3, 0.9)
	glContext.Clear(webgl.COLOR_BUFFER_BIT)
}

// ----------------------------------------------------------------------------
func setupViewport() {
	glContext.Enable(webgl.DEPTH_TEST)
	glContext.Viewport(0, 0, clientWidth, clientHeight)
}
