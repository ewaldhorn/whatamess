package main

import (
	"errors"
	"fmt"
)

// ----------------------------------------------------------------------------
func main() {
	err := triggerError() // trigger an error

	var anError *CustomError

	// check if it's our custom error
	if errors.As(err, &anError) {
		fmt.Println("concrete error type")
		fmt.Println(anError.ErrorFunction())
	}

	// now check if it's our interface
	var customerErrorInterface SomeInterface
	if errors.As(err, &customerErrorInterface) {
		fmt.Println("it's also the interface hey")
		fmt.Println(customerErrorInterface.ErrorFunction())
	}
}

// ----------------------------------------------------------------------------
func triggerError() error {
	err := CustomError{}
	return err
}

// ----------------------------------------------------------------------------
type CustomError struct{}

// ----------------------------------------------------------------------------
func (c CustomError) Error() string {
	return "custom error"
}

// ----------------------------------------------------------------------------
func (c *CustomError) ErrorFunction() string {
	return "this is a returned string"
}

// ----------------------------------------------------------------------------
type SomeInterface interface {
	ErrorFunction() string
}
