package main

import "fmt"

type age int       //underlying type is int
type oldAge age    //underlying type is int
type verOldAge age //underlying type is int

type speed uint
type km float64
type mile float64

func main() {
	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	// can't work, the types are not the same: speed vs uint
	// var x1 uint = 10
	// fmt.Println(s2 - x1)
	var x uint
	x = uint(s1)
	fmt.Println("Casting to uint works:", x)

	var s3 speed = speed(x)
	fmt.Println("s3 is:", s3)

	var parisToLondon km = 465
	var distanceInMiles mile
	distanceInMiles = mile(parisToLondon) / 0.621
	fmt.Println("Paris to London is", parisToLondon, "in km or", distanceInMiles, "in miles")
}
