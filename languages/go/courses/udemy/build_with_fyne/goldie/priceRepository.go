package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const DataURL = "https://data-asg.goldprice.org/dbXRates"
const ChartURL = "https://goldprice.org/charts/gold_3d_b_o_usd_x.png"

var currency = "USD"

type Gold struct {
	Prices []Price `json:"items"`
	Client *http.Client
}

type Price struct {
	Currency      string    `json:"currency"`
	Price         float64   `json:"xauPrice"`
	Change        float64   `json:"chgXau"`
	PreviousClose float64   `json:"xauClose"`
	Time          time.Time `json:"-"`
}

func (g *Gold) getPrices() (*Price, error) {
	if g.Client == nil {
		// no client, can't help ya yet!
		g.Client = &http.Client{}
	}

	client := g.Client
	url := fmt.Sprintf("%s/%s", DataURL, currency)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		myApp.ErrorLog.Println("error creating request", err)
		return nil, err
	}

	rsp, err := client.Do(req)

	if err != nil {
		log.Println("error retrieving gold price", err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("unable to close response body", err)
		}
	}(rsp.Body)

	body, err := io.ReadAll(rsp.Body)

	if err != nil {
		log.Println("unable to read response body", err)
		return nil, err
	}

	gold := Gold{}
	var previous, current, change float64
	err = json.Unmarshal(body, &gold)
	if err != nil {
		log.Println("error unmarshalling json", err)
		return nil, err
	}

	previous, current, change = gold.Prices[0].PreviousClose, gold.Prices[0].Price, gold.Prices[0].Change

	var currentInfo = Price{
		Currency:      currency,
		Price:         current,
		Change:        change,
		PreviousClose: previous,
		Time:          time.Now(),
	}

	return &currentInfo, nil
}

func (cfg *Config) getPriceChart() *canvas.Image {
	apiURL := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(currency))
	var img *canvas.Image

	err := cfg.downloadFile(apiURL, "gold.png")
	if err != nil {
		log.Println("error retrieving price chart", err)
		// use default image
		img = canvas.NewImageFromResource(resourceUnreachablePng)
	} else {
		img = canvas.NewImageFromFile("gold.png")
	}

	img.SetMinSize(fyne.Size{
		Width:  780,
		Height: 500,
	})

	img.FillMode = canvas.ImageFillOriginal
	return img
}
