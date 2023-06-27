package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"
)

var testApp Config

var jsonToReturn = `{"ts":1687874756262,"tsj":1687874756098,"date":"Jun 27th 2023, 10:05:56 am NY","items":[{"curr":"USD","xauPrice":1234.5678,"xagPrice":22.8427,"chgXau":-1.9225,"chgXag":0.0807,"pcXau":-0.1,"pcXag":0.3545,"xauClose":1922.505,"xagClose":22.762}]}`

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGold_Basics(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})
	g := Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.getPrices()

	if err != nil {
		t.Error(err)
	}

	if p.Price != 1234.5678 {
		t.Error("wrong price returned:", p.Price)
	}
}

// Create a custom client and transport function that'll return some static JSON
type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

//func TestGold_getPrices(t *testing.T) {
//	type fields struct {
//		Prices []Price
//		Client *http.Client
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		want    *Price
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			g := &Gold{
//				Prices: tt.fields.Prices,
//				Client: tt.fields.Client,
//			}
//			got, err := g.getPrices()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("getPrices() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("getPrices() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
