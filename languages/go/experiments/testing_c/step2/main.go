package main

/*
#cgo CFLAGS: -g
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>

void responder(void) {
   printf("This is C. Read you five by five.\n");
}

void sayHelloTo(char* person) {
   printf("Oh, howdy %s!\n", person);
}

char* toUpper(char* s) {
 for (char *ptr=s; *ptr; ptr++) *ptr=toupper(*ptr);
 return s;
}

int populate(int quantity, const char *item, char *out) {
   return sprintf(out, "We have %d %s available.", quantity, item);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// ======================================================================= main
func main() {
	fmt.Println("Attempting to call C.")
	C.responder()
	C.sayHelloTo(C.CString("Roger"))
	fmt.Println("abcd becomes:", C.GoString(C.toUpper(C.CString("abcd"))))

	// unsafe operations follow

	// create parameters
	productName := C.CString("Groovy Cheese Straws")
	defer C.free(unsafe.Pointer(productName))

	stockCount := C.int(11)

	// create buffer
	bufferPtr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(bufferPtr))

	// make the call, passing the parameters and the buffer
	size := C.populate(stockCount, productName, (*C.char)(bufferPtr))

	// copy the buffer into a Go memory-managed slice
	goSlice := C.GoBytes(bufferPtr, size)

	// print the slice we created from the C buffer
	fmt.Println(string(goSlice))
}
