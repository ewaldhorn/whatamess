package main

import (
	"fmt"

	"github.com/carlmjohnson/versioninfo"
)

func main() {
	fmt.Println(versioninfo.Version)
	fmt.Println(versioninfo.Short())
}
