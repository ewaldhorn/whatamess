package main

import (
	"net/url"
	"strings"
)

// ----------------------------------------------------------------------------
type errors map[string][]string

// ----------------------------------------------------------------------------
func (e errors) Get(field string) string {
	errorSlice := e[field]

	if len(errorSlice) == 0 {
		return ""
	}

	return errorSlice[0]
}

// ----------------------------------------------------------------------------
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// ----------------------------------------------------------------------------
type Form struct {
	Data   url.Values
	Errors errors
}

// ----------------------------------------------------------------------------
func NewForm(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: map[string][]string{},
	}
}

// ----------------------------------------------------------------------------
// checks whether a field exists in a form
func (f *Form) Has(field string) bool {
	return f.Data.Get(field) != ""
}

// ----------------------------------------------------------------------------
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field can not be blank")
		}
	}
}

// ----------------------------------------------------------------------------
func (f *Form) Check(ok bool, key, message string) {
	if !ok {
		f.Errors.Add(key, message)
	}
}

// ----------------------------------------------------------------------------
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
