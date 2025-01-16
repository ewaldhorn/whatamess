package main

import (
	"log"

	"github.com/unidoc/unipdf/v3/creator"
)

var pdfCreator *creator.Creator

// ----------------------------------------------------------------------------
func addParagraph(x, y float64, title string) *creator.Paragraph {
	paragraph := pdfCreator.NewParagraph(title)
	paragraph.SetPos(x, y)

	return paragraph
}

// ----------------------------------------------------------------------------
func renderParagraph(paragraph *creator.Paragraph) {
	pdfCreator.Draw(paragraph)
}

func main() {
	pdfCreator = creator.New()
	pdfCreator.NewPage()

	// add a new paragraph to the page
	paragraph := addParagraph(50, 70, "About")
	paragraph.SetText("This is the About section.")
	renderParagraph(paragraph)

	if err := pdfCreator.WriteToFile("generated.pdf"); err != nil {
		log.Fatalf("PDF creation error: %v", err)
	}
}
