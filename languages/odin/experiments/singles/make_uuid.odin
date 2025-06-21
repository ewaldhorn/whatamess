package main

import "core:crypto"
import "core:encoding/uuid"
import "core:fmt"

main :: proc() {
	my_uuid: uuid.Identifier

	{
		// This scope will have a CSPRNG.
		context.random_generator = crypto.random_generator()
		my_uuid = uuid.generate_v7()
	}

	// Back to the default random number generator.
	fmt.println("The uuid itself is   :", my_uuid)
	fmt.printf("Generated UUID string: %s\n", uuid.to_string(my_uuid))
}
