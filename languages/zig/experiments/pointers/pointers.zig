const std = @import("std");
const expect = std.testing.expect;

// accepts a pointer to a u8 and adds 1 to the unsigned 8 bit number
fn increment(num: *u8) void {
    num.* += 1;
}

test "pointers" {
    var x: u8 = 1;
    increment(&x);
    increment(&x);
    try expect(x == 3);
}

// can't set a pointer to 0/null
// test will fail
// test "naughty pointer" {
//     const x: u16 = 0;
//     const y: *u8 = @ptrFromInt(x);
//     _ = y;
// }

// can't mutate a const pointer
// const x: u8 = 1; creates a const pointer, can't change that value
// test will fail
// test "const pointers" {
//     const x: u8 = 1;
//     const y = &x;
//     y.* += 1;
// }

// usize and isize are given as unsigned and signed integers which are the same size as pointers.
test "usize" {
    try expect(@sizeOf(usize) == @sizeOf(*u8));
    try expect(@sizeOf(isize) == @sizeOf(*u8));
}