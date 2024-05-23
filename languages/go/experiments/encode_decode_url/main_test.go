package main

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_singleString(t *testing.T) {
	result := escapeString("https://www.nofuss.co.za/?this=that")
	expected := "https%3A%2F%2Fwww.nofuss.co.za%2F%3Fthis%3Dthat"
	if result != expected {
		t.Fatalf("expected [%s], got [%s]\n", expected, result)
	}
}

func Test_decodingOfURL(t *testing.T) {
	result, err := url.Parse("https%3A%2F%2Fwww.nofuss.co.za%2F%3Fthis%3Dthat")

	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	fmt.Println(result)

}

func Test_paramEncoding(t *testing.T) {
	result := escapeQueryParameters("$4", "&nope!!", "just 😀")
	expected := "param_1=%244&param_2=%26nope%21%21&param_3=just+%F0%9F%98%80"
	if result != expected {
		t.Fatalf("expected [%s], got [%s]\n", expected, result)
	}
}

func Test_allInOne(t *testing.T) {
	result := allInOne("https", "hereorthere.com", "/version")
	expected := "https://hereorthere.com/version"

	if result != expected {
		t.Fatalf("expected [%s], got [%s]\n", expected, result)
	}
}
