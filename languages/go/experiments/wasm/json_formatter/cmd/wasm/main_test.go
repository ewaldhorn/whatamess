package main

import "testing"

func TestJSONFormatter(t *testing.T) {
	t.Run("empty JSON", func(t *testing.T){
		got, err := prettyJson("{}")
		want := "{}"

		if err != nil {
			t.Errorf("got %q wanted no error", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}) 

	t.Run("simple JSON", func(t *testing.T){
		got, err := prettyJson(`{"website":"yahoo.com"}`)
		want := `{
  "website": "yahoo.com"
}`

		if err != nil {
			t.Errorf("got %q wanted no error", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}) 

}