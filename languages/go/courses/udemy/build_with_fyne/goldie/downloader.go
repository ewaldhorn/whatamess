package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func (cfg *Config) downloadFile(URL, filename string) error {
	// get response bytes from calling a url
	rsp, err := cfg.HTTPClient.Get(ChartURL)

	if err != nil {
		log.Println("error getting url", err)
		return err
	}

	if rsp.StatusCode != 200 {
		log.Println("error downloading chart: ", rsp.StatusCode)
		return errors.New("error downloading chart")
	}

	bytesRead, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Println("error reading response body", err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("error close response body", err)
		}
	}(rsp.Body)

	img, _, err := image.Decode(bytes.NewReader(bytesRead))
	if err != nil {
		log.Println("error decoding image", err)
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s", filename))
	if err != nil {
		log.Println("unable to create output file", err)
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		log.Println("error writing png", err)
		return err
	}

	return nil
}
