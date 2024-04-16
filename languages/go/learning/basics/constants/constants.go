package main

func main() {
	const daysInWeek int = 7
	const averageWeeksPerMonth = 4
	const averageDaysPerMonth = daysInWeek * averageWeeksPerMonth

	println("We have an average of", averageDaysPerMonth, "days per month.")

	const (
		c1 = iota
		c2 = iota
		c3 = iota
		c4 = iota + 5
		_
		_
		c5
		c6
		_
		_
		c7 = iota * 2
		c8
		c9
	)

	println("c1:", c1)
	println("c2:", c2)
	println("c3:", c3)
	println("c4:", c4)
	println("c5:", c5)
	println("c6:", c6)
	println("c7:", c7)
	println("c8:", c8)
	println("c9:", c9)
}
