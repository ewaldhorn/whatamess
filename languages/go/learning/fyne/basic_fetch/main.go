package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()

	window := application.NewWindow("Fyne Fetch")
	window.Resize(fyne.NewSize(800.0, 600.0))
	window.CenterOnScreen()

	label := widget.NewLabel("New Content")
	vBox := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), label, layout.NewSpacer())
	window.SetContent(vBox)

	client = &http.Client{Timeout: time.Second * 10}
	fact, err := randomFact()

	if err != nil {
		label.SetText("Error retrieving a random fact!")
	} else {
		label.SetText(fact)
	}

	window.ShowAndRun()
}

var client *http.Client

func randomFact() (string, error) {
	rsp, err := client.Get("https://uselessfacts.jsph.pl/random.json")

	if err != nil {
		log.Println("Error retrieving fact:", err)
		return "<< Error >>", err
	}

	defer rsp.Body.Close()

	fact, err := decodeRandomFact(rsp.Body)

	if err != nil {
		log.Println("Unable to parse random fact: ", err)
		return "", err
	}

	return fact, nil
}

func decodeRandomFact(jsonString io.Reader) (string, error) {
	body := struct {
		Text string `json:"Text"`
	}{}

	err := json.NewDecoder(jsonString).Decode(&body)
	if err != nil {
		log.Println("Could not parse JSON:", err)
		return "", err
	}

	return body.Text, nil
}
