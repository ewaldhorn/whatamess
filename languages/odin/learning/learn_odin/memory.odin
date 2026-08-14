package main

// I first wrote the code, then, for giggles, hit up Qwen3.6 to ask it
// to add documentation. I asked it to be verbose and it did not hold
// back. Pretty impressive considering it's a local model!

import "core:fmt"

// User is a simple struct used to demonstrate dynamic allocation of
// composite (record) types in Odin. It holds an integer identifier and
// a string name.
//
// In Odin, records are value types by default — copying a `User` copies
// every field. When placed inside a dynamic array (`[dynamic]User`), each
// element lives on the heap and is managed by the allocator passed to
// `make`.
User :: struct {
	id:   int,
	name: string,
}

// ------------------------------------------------------------------------------------------------
// main demonstrates four fundamental memory-management concepts in Odin:
//
//  1. **Fixed arrays** — compile-time sized, stack-allocated.
//  2. **Slices** — runtime-sized views into a contiguous memory region.
//  3. **Dynamic arrays** — heap-allocated, resizable collections.
//  4. **Scratch buffers** — temporary, reusable memory for short-lived work.
//
// Understanding the distinction between these types is essential for
// writing correct and efficient Odin programs, because each has different
// lifetime, allocation, and mutation semantics.
//
// See also:
//   - https://odin-lang.org/docs/intro/
//   - https://github.com/odin-lang/Odin/blob/master/docs/README.md
// ------------------------------------------------------------------------------------------------
main :: proc() {
	// ──────────────────────────────────────────────────────────────────
	// 1. Fixed array
	//
	// A fixed array has a size known at compile time. The compiler places
	// the entire array on the **stack**, which means:
	//
	//   • Allocation is essentially free — the stack pointer moves.
	//   • The memory is automatically reclaimed when the function returns
	//     (no explicit deallocation needed).
	//   • The size is part of the type: `[3]int` is a *different type*
	//     from `[4]int`. They are not interchangeable.
	//
	// Fixed arrays are ideal for small, known-size collections such as
	// lookup tables, fixed buffers, or coordinate tuples.
	//
	// NOTE: The original comment said "will now change size" — that is
	// misleading. A fixed array's size is immutable by definition. If
	// you need to change the size, you must create a new array and copy
	// the elements. See the dynamic array below for a resizable alternative.
	fixed_scores: [3]int = {90, 85, 95}

	// ──────────────────────────────────────────────────────────────────
	// 2. Slice (a view into memory)
	//
	// A slice does **not** own its data. Instead, it is a small header
	// (pointer + length + capacity) that refers to a contiguous region of
	// memory — in this case, the fixed array above.
	//
	//   `fixed_scores[0:2]` creates a slice that:
	//     • Starts at index 0 of `fixed_scores`.
	//     • Extends 2 elements forward (indices 0 and 1).
	//     • Points to the *same underlying memory* as `fixed_scores`.
	//
	// Key properties of slices:
	//
	//   • **Shared ownership** — modifying an element through the slice
	//     also modifies the original array, because they share the same
	//     backing memory.
	//
	//   • **Zero-copy** — no data is copied when a slice is created.
	//     This makes slicing very cheap.
	//
	//   • **Bounded** — the slice cannot access elements outside its
	//     length. Attempting to do so is a panic at runtime.
	//
	//   • **Not array-sized** — `[]int` is a completely different type
	//     from `[3]int`. Slices are dynamically sized; arrays are not.
	//
	// Slices are the workhorse of Odin (and Go) data handling. They give
	// you the flexibility of dynamic arrays without the allocation cost,
	// as long as the underlying data already exists.
	score_view: []int = fixed_scores[0:2] // gives {90, 85}

	// ──────────────────────────────────────────────────────────────────
	// 3. Dynamic array (heap-allocated, resizable)
	//
	// A dynamic array is a heap-allocated, resizable collection. It is
	// created with `make` and manages its own memory via an **allocator**.
	//
	//   `make([dynamic]User, context.allocator)` allocates an empty
	//   dynamic array of `User` records, using the global default
	//   allocator (`context.allocator`).
	//
	// Key properties:
	//
	//   • **Heap allocation** — the backing store lives on the heap, not
	//     the stack. This means the data persists beyond the current
	//     function call (if a reference is passed out).
	//
	//   • **Resizable** — elements can be added with `append(&arr,
	//     value)`. When the array fills its current capacity, `append`
	//     automatically allocates a larger backing store and copies the
	//     elements over. This reallocation is amortized O(1).
	//
	//   • **Manual cleanup** — unlike stack-allocated arrays, dynamic
	//     arrays must be explicitly freed. `defer delete(users)` ensures
	//     the memory is reclaimed when `main` returns, preventing a
	//     memory leak.
	//
	//   • **Pass by reference** — `append` takes a pointer (`&users`)
	//     because it may modify the slice header (pointer, length,
	//     capacity) in place.
	//
	// Dynamic arrays are the go-to choice when the number of elements is
	// not known at compile time, or when the collection needs to grow
	// and shrink over time.
	users := make([dynamic]User, context.allocator)
	defer delete(users)

	append(&users, User{id = 1, name = "Allsop"})
	append(&users, User{2, "Jones"})

	// ──────────────────────────────────────────────────────────────────
	// 4. Scratch buffer (temporary memory)
	//
	// A scratch buffer is a pre-allocated block of memory intended for
	// short-lived, temporary use. It is created with `context.temp_allocator`,
	// which is designed for allocations that will be discarded en masse
	// rather than individually freed.
	//
	//   `make([]u8, 512, context.temp_allocator)` creates a 512-byte
	//   buffer of `u8` (unsigned 8-bit integer / byte) values, suitable
	//   for temporary work such as:
	//
	//     • Building a string incrementally.
	//     • Parsing input into a mutable buffer.
	//     • Acting as a temporary workspace for computations.
	//
	// Key properties:
	//
	//   • **Reusable** — scratch buffers are typically allocated once at
	//     startup and reused throughout the program's lifetime, avoiding
	//     repeated allocation overhead.
	//
	//   • **Capacity-focused** — the `len` gives the usable size, while
	//     the underlying capacity may be larger. You can grow or shrink
	//     the usable portion with `setlen`.
	//
	//   • **No individual free** — you do not call `delete` on a scratch
	//     buffer. Its memory is managed by the temp allocator, which
	//     typically reclaims it all at program exit or at explicit
	//     checkpoint/rollback points.
	//
	// Scratch buffers are a performance optimization for scenarios where
	// you need mutable temporary storage but want to avoid the overhead
	// of repeated small allocations.
	scratch_buffer := make([]u8, 512, context.temp_allocator)

	// ──────────────────────────────────────────────────────────────────
	// Output: print each memory concept to verify correctness.
	//
	// `fmt.println` handles formatting for all four types:
	//   • Arrays → `{90 85 95}`
	//   • Slices → `{90 85}`
	//   • Dynamic arrays → `{User{id:1, name:"Allsop"} User{id:2, name:"Jones"}}`
	//   • Scratch buffer → `512` (its length)
	//
	// `fmt.printf` with `%v` prints the length of the scratch buffer.
	fmt.println("Fixed scores  :", fixed_scores)
	fmt.println("Slice view    :", score_view)
	fmt.println("Dynamic users :", users)
	fmt.printf("Scratch buffer : %v bytes\n", len(scratch_buffer))
}
