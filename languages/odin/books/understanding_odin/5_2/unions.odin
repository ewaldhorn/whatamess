package unions

import "core:fmt"

// ----------------------------------------------------------------------------
My_Union :: union {
	f32,
	int,
	Person_Data,
}

// ----------------------------------------------------------------------------
Person_Data :: struct {
	health: int,
	age:    int,
}

// ----------------------------------------------------------------------------
val: My_Union = int(12)

// ----------------------------------------------------------------------------
main :: proc() {
	fmt.printfln("We have %v", val)

	val = Person_Data {
		health = 76,
		age    = 14,
	}

	fmt.printfln("We have %v", val)

	// switching over a union on the type
	switch v in val {
	case int:
		fmt.println("An int")

	case f32:
		fmt.println("A float")

	case Person_Data:
		fmt.println("Person_Data")
	}

	// check if a union is of a certain type
	f32_val, f32_val_ok := val.(f32)

	if f32_val_ok {
		fmt.println("We got a float")
	} else {
		fmt.println("Nope, no float here")
	}

}
