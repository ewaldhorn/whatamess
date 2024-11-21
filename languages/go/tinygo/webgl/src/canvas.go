package main

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/core/js"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/graphics/webgl"
	"github.com/gowebapi/webapi/html/canvas"
)

// ----------------------------------------------------------------------------
var (
	glContext                 *webgl.RenderingContext
	vertexBuffer, indexBuffer *webgl.Buffer
	indicesCount              int
	shaderProgram             *webgl.Program
	clientWidth, clientHeight int
)

// ----------------------------------------------------------------------------
func createCanvas() {
	doc, app := getDocAndAppHandles()

	rawCanvasElement := getRawCanvasElement()
	app.AppendChild(&rawCanvasElement.Node)

	canvasHTML := setupHTMLCanvasElement(doc, rawCanvasElement)
	glContext = webgl.RenderingContextFromWrapper(canvasHTML.GetContext("webgl", nil))

	vertexBuffer, indexBuffer, indicesCount = createBuffers(glContext)

	shaderProgram = setupShaders(glContext)
}

// ----------------------------------------------------------------------------
func getDocAndAppHandles() (*webapi.Document, *dom.Element) {
	tmpDoc := webapi.GetWindow().Document()
	return tmpDoc, tmpDoc.GetElementById("app")
}

// ----------------------------------------------------------------------------
func getRawCanvasElement() *dom.Element {
	tmpElement := webapi.GetWindow().Document().CreateElement("canvas", &webapi.Union{Value: js.ValueOf("dom.Node")})
	tmpElement.SetId("canvas")
	return tmpElement
}

// ----------------------------------------------------------------------------
func setupHTMLCanvasElement(doc *webapi.Document, canvasElement *dom.Element) *canvas.HTMLCanvasElement {
	tmpCanvasHTMLElement := canvas.HTMLCanvasElementFromWrapper(canvasElement)

	body := doc.GetElementById("body")

	clientWidth = body.ClientWidth()
	clientHeight = body.ClientHeight()

	tmpCanvasHTMLElement.SetWidth(uint(clientWidth))
	tmpCanvasHTMLElement.SetHeight(uint(clientHeight))

	return tmpCanvasHTMLElement
}
