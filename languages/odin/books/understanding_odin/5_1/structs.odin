package structs

import "core:fmt"

Point :: struct {
	x, y: f32,
}

// ----------------------------------------------------------------------------
Rectangle :: struct {
	using point: Point,
	width:       f32,
	height:      f32,
}

// ----------------------------------------------------------------------------
add_points :: proc(p1, p2 : Point) -> Point {
	return Point {
		x = p1.x + p2.x,
		y = p1.y + p2.y
	}
}

// ----------------------------------------------------------------------------
multiply_points::proc(p1,p2:Point)->Point{
	return Point{x=p1.x*p2.x, y=p1.y*p2.y}
}

// ----------------------------------------------------------------------------
main :: proc() {
	myRect := Rectangle {
		x      = 0.0,
		y      = 0.0,
		width  = 100.0,
		height = 100.0,
	}

	otherRect: Rectangle = {
		width  = 50.0,
		height = 50.0,
	}

	p1 := Point{x=1, y=2}
	p2 := Point{x=2, y=1}
	p3 := add_points(p1,p2)

	fmt.println("P3 is:", p3)

	p4 := multiply_points(p1,p2)
	fmt.println("P4 is:", p4)
}
