const expect = @import("std").testing.expect;
const meta = @import("std").meta;

// Zig provides vector types for SIMD. These are not to be conflated with vectors in a mathematical sense, 
// or vectors like C++'s std::vector (for this, see "Arraylist" in chapter 2). Vectors may be created 
// using the @Type built-in we used earlier, and std.meta.Vector provides a shorthand for this.

// Vectors can only have child types of booleans, integers, floats and pointers.

// Operations between vectors with the same child type and length can take place. These operations are 
// performed on each of the values in the vector.std.meta.eql is used here to check for equality between 
// two vectors (also useful for other types like structs).

test "vector add" {
    const x: @Vector(4, f32) = .{ 1, -10, 20, -1 };
    const y: @Vector(4, f32) = .{ 2, 10, 0, 1 };
    const z = x + y;
    try expect(meta.eql(z, @Vector(4, f32){ 3, 0, 20, 0 }));
}

test "vector indexing" {
    const x: @Vector(4, u8) = .{ 255, 0, 255, 0 };
    try expect(x[0] == 255);
}

test "vector * scalar" {
    const x: @Vector(3, f32) = .{ 12.5, 37.5, 2.5 };
    const y = x * @as(@Vector(3, f32), @splat(2));
    try expect(meta.eql(y, @Vector(3, f32){ 25, 75, 5 }));
}

test "vector looping" {
    const x = @Vector(4, u8){ 255, 0, 255, 0 };
    const sum = blk: {
        var tmp: u10 = 0;
        var i: u8 = 0;
        while (i < 4) : (i += 1) tmp += x[i];
        break :blk tmp;
    };
    try expect(sum == 510);
}

const arr: [4]f32 = @Vector(4, f32){ 1, 2, 3, 4 };