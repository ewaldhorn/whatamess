package structs

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
}
