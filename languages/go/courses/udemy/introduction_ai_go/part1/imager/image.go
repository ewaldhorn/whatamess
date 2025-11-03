package imager

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"mazes/globals"
	"mazes/maze"
	"mazes/node"
	"mazes/point"
	"mazes/wall"
	"os"

	"github.com/StephaneBunel/bresenham"
	"github.com/kmicki/apng"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// ------------------------------------------------------------------------------------------------
const cellSize = 60

// ------------------------------------------------------------------------------------------------
var (
	green     = color.RGBA{G: 255, A: 255}
	darkGreen = color.RGBA{R: 1, G: 100, B: 32, A: 255}
	red       = color.RGBA{R: 255, A: 255}
	yellow    = color.RGBA{R: 255, G: 255, B: 101, A: 255}
	gray      = color.RGBA{R: 125, G: 125, B: 125, A: 125}
	orange    = color.RGBA{R: 255, G: 140, B: 25, A: 255}
	blue      = color.RGBA{R: 14, G: 118, B: 173, A: 255}
)

// ------------------------------------------------------------------------------------------------
func RenderMazeToDisk(m *maze.Maze, filename ...string) {
	fmt.Printf("Generating image: `%s`\n", filename[0])

	// calculate image dimensions
	width := cellSize * (m.Width - 1)
	height := cellSize * m.Height

	// ensure filename is valid
	var outFile = "image.png"
	if len(filename) > 0 {
		outFile = filename[0]
	}

	// set output dimensions
	upLeft := image.Point{}
	bottomRight := image.Point{X: width, Y: height}

	// create output image
	img := image.NewRGBA(image.Rectangle{Min: image.Point(upLeft), Max: image.Point(bottomRight)})
	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.Black}, image.Point{}, draw.Src)

	// draw the actual blocks
	drawBlocks(img, m)

	// draw the grid
	drawGrid(img, m)

	f, _ := os.Create(outFile)
	_ = png.Encode(f, img)
}

// ------------------------------------------------------------------------------------------------
func RenderAnimatedMazeToDisk() {
	output := "./animation.png"
	files, _ := os.ReadDir(globals.TemporaryDirectory)
	var images []string
	var delays []int

	for _, file := range files {
		images = append(images, fmt.Sprintf("%s%s", globals.TemporaryDirectory, file.Name()))
		delays = append(delays, 30)
	}

	images = append(images, "./image.png")

	a := apng.APNG{
		Frames: make([]apng.Frame, len(images)),
	}
	out, _ := os.Create(output)
	defer out.Close()

	for i, s := range images {
		in, err := os.Open(s)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		defer in.Close()

		m, err := png.Decode(in)
		if err != nil {
			continue
		}
		a.Frames[i].Image = m
	}
	err := apng.Encode(out, a)
	if err != nil {
		log.Println(err)
	}

}

// ------------------------------------------------------------------------------------------------
func drawBlocks(img *image.RGBA, m *maze.Maze) {
	// Precompute expensive lookups
	solutionSet := make(map[point.Point]bool)
	exploredSet := make(map[point.Point]bool)

	for _, p := range m.Solution.Cells {
		solutionSet[p] = true
	}

	for _, p := range m.Explored {
		exploredSet[p] = true
	}

	// now check and render each block
	for i, row := range m.Walls {
		for j, col := range row {
			p := point.Point{X: j, Y: i}
			x, y := j*cellSize, i*cellSize

			var fillColor color.Color

			switch {
			case col.IsSolid:
				fillColor = color.Black
			case p.X == m.Start.X && p.Y == m.Start.Y:
				fillColor = darkGreen
			case p.X == m.End.X && p.Y == m.End.Y:
				fillColor = red
			case solutionSet[p]:
				fillColor = green
			case col.State == m.CurrentNode.State:
				fillColor = orange
			case exploredSet[p]:
				fillColor = yellow
			default:
				fillColor = color.White
			}

			drawSquare(col, p, img, fillColor, cellSize, x, y, mustDrawCost(m.SearchType), m.Start)
		}
	}
}

// -------------------------------------------------------------------------------------------------
func mustDrawCost(searchType int) bool {
	switch searchType {
	case globals.DIJKSTRA:
		return true
	default:
		return false
	}
}

// ------------------------------------------------------------------------------------------------
func drawGrid(img *image.RGBA, m *maze.Maze) {
	for y, _ := range m.Walls {
		bresenham.DrawLine(img, 0, y*cellSize, m.Width*cellSize, y*cellSize, gray)
	}

	for x, _ := range m.Walls[0] {
		bresenham.DrawLine(img, x*cellSize, 0, x*cellSize, m.Height*cellSize, gray)
	}

	bresenham.DrawLine(img, m.Width*cellSize-cellSize-1, 0, m.Width*cellSize-cellSize-1, m.Height*cellSize, gray)
	bresenham.DrawLine(img, 0, m.Height*cellSize-1, m.Width*cellSize, m.Height*cellSize-1, gray)
}

// ------------------------------------------------------------------------------------------------
func drawSquare(
	col wall.Wall,
	p point.Point,
	img *image.RGBA,
	c color.Color,
	size, x, y int,
	drawCost bool,
	target point.Point,
) {
	patch := image.NewRGBA(image.Rect(0, 0, size, size))

	draw.Draw(patch, patch.Bounds(), &image.Uniform{C: c}, image.Point{}, draw.Src)

	if !col.IsSolid {
		if drawCost {
			printManhattanCost(p, color.Black, patch, target)
		}
		printLocation(p, color.Black, patch)
	}

	draw.Draw(img, image.Rect(x, y, x+size, y+size), patch, image.Point{}, draw.Src)
}

// ------------------------------------------------------------------------------------------------
func printLocation(p point.Point, c color.Color, patch *image.RGBA) {
	point := fixed.Point26_6{X: fixed.I(6), Y: fixed.I(40)}
	d := &font.Drawer{Dst: patch, Src: image.NewUniform(c), Face: basicfont.Face7x13, Dot: point}
	d.DrawString(fmt.Sprintf("[%d %d]", p.X, p.Y))
}

// -------------------------------------------------------------------------------------------------
func printManhattanCost(p point.Point, c color.Color, patch *image.RGBA, target point.Point) {
	point := fixed.Point26_6{X: fixed.I(6), Y: fixed.I(17)}
	d := &font.Drawer{Dst: patch, Src: image.NewUniform(c), Face: basicfont.Face7x13, Dot: point}
	n := node.Node{State: p}
	d.DrawString(fmt.Sprintf("%d", n.ManhattanDistance(target)))
}
