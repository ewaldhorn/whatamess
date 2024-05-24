package main

import (
	"testing"
)

// ------------------------------------------------------------------------- Test_accessMapElements
func Test_accessMapElements(t *testing.T) {
	myMap := map[string]int{
		"acceleration": 11,
		"performance":  5,
		"cost":         100,
	}

	if myMap["cost"] != 100 {
		t.Error("cost was not as expected!")
	}

	// we should not find this
	_, found := myMap["doofus"]
	if found == true {
		t.Error("unexpected!")
	}
}

// ---------------------------------------------------------------------------- Test_addingElements
func Test_addingElements(t *testing.T) {
	myMap := map[string]int{
		"acceleration": 9,
	}

	if len(myMap) != 1 {
		t.Fatal("myMap should only have one element")
	}

	myMap["cost"] = 900

	if len(myMap) != 2 {
		t.Fatal("myMap should now have two elements")
	}
}

// ---------------------------------------------------------------------------- Test_removeElements
func Test_removeElements(t *testing.T) {
	myMap := map[string]int{
		"acceleration": 11,
		"performance":  5,
		"cost":         100,
	}

	if len(myMap) != 3 {
		t.Fatal("myMap should have three elements")
	}

	delete(myMap, "cost")

	if len(myMap) != 2 {
		t.Fatal("myMap should have two elements")
	}

	cost, found := myMap["cost"]

	if found {
		t.Fatal("did not expect cost:", cost)
	}

	delete(myMap, "this key doesn't exist")
	if len(myMap) != 2 {
		t.Fatal("myMap should have two elements")
	}
}
