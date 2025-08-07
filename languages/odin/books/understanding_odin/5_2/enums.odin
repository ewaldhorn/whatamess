package enums

import "core:fmt"

// ----------------------------------------------------------------------------
Computer_Type :: enum {
	Laptop, // value 0
	Desktop, // value 1
	Mainframe, // value 2
}

// ----------------------------------------------------------------------------
Rig_Type :: enum {
	Laptop = 0, // value 0
	Desktop, // value 1
	Mainframe = 5, // value 5
	Broken, // value 6
	Gamer_Rig, // value 7
}

// ----------------------------------------------------------------------------
main :: proc() {
	item: Computer_Type

	item = .Desktop

	switch item {
	case .Laptop:
		fmt.println("Laptop")
	case .Desktop:
		fmt.println("Desktop")
	case .Mainframe:
		fmt.println("Mainframe")
	}
}
