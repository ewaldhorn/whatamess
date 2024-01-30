const std = @import("std");

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const expect = std.testing.expect;

test "basic array operations" {
    const my_array = [_]i32{ 0, 1, 2, 3, 4, 5, 6 };
    const my_double_array = my_array ** 2;

    try expect(my_array.len == 7);
    try expect(my_double_array.len == 14);
}

test "undefined array operations" {
    var my_array : [3]u8 = undefined;

    my_array[0] = 1;

    try expect(my_array.len == 3);
    try expect(my_array[0] == 1);
}

test "multi-dimensional array operations" {
    const my_array = [_][3]u8 {.{1,2,3}, .{4,5,6}};

    try expect(my_array.len == 2);
    try expect(my_array[0].len == 3);
}

test "array sentinels" {
    const my_array : [2:'a']u8 = .{11, 12};

    try expect(my_array.len == 2);
    try expect(my_array[my_array.len] == 'a');
}

test "concatenate arrays" {
    const a1 = [_]u8{0,1,2,3};
    const b1 = [_]u8{4,5,6,7,8,9};
    const combined = a1 ++ b1;

    try expect(combined.len == (a1.len + b1.len));
}
