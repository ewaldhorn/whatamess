package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// ----------------------------------------------------------------------------
func TestForm_Has(t *testing.T) {
	form := NewForm(nil)

	has := form.Has("not_expected_to_be_found")

	if has {
		t.Error("form shows an expected field, expected to be empty")
	}

	postedData := url.Values{}
	postedData.Add("field", "value")

	form = NewForm(postedData)

	has = form.Has("field")
	if !has {
		t.Error("expected to find 'field', failed")
	}
}

// ----------------------------------------------------------------------------
func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/some_url", nil)
	form := NewForm(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows fields, but it should not be")
	}

	postedData := url.Values{}
	postedData.Add("left", "a")
	postedData.Add("middle", "b")
	postedData.Add("right", "c")

	r, _ = http.NewRequest("POST", "/some_url", nil)
	r.PostForm = postedData

	form = NewForm(r.PostForm)
	form.Required("left", "middle", "right")
	if !form.Valid() {
		t.Error("post does not have required fields, but it should have")
	}
}

// ----------------------------------------------------------------------------
func TestForm_Check(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")

	if form.Valid() {
		t.Error("Valid() returns false, and it should be true")
	}
}

// ----------------------------------------------------------------------------
func TestForm_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")

	s := form.Errors.Get("password")
	if len(s) == 0 {
		t.Error("should have an error returned from Get, but did not")
	}

	s = form.Errors.Get("whatever")
	if len(s) != 0 {
		t.Error("should not have an error returned from Get, but did")
	}
}
