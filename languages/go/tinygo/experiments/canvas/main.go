package main

import (
 "tinygo.org/x/tinycanvas"
)

func main() {
 // Create a new TinyCanvas instance
 tc := tinycanvas.New()

 // Set the canvas dimensions
 tc.SetDimensions(400, 400)

 // Clear the canvas with a white background
 tc.Clear(tinycanvas.ColorWhite)

 // Set the fill color to red
 tc.SetFillColor(tinycanvas.ColorRed)

 // Draw a rectangle at position (100, 100) with dimensions 200x200
 tc.FillRect(100, 100, 200, 200)

 // Save the canvas to an HTML file
 tc.Save("rectangle.html")
}
