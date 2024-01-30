const std = @import("std");

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const expect = std.testing.expect;
const allocator = std.testing.allocator;

test "basic array list operations" {
    var my_list = std.ArrayList(u8).init(allocator);
    defer my_list.deinit();

    try expect(my_list.items.len == 0);

    try my_list.append(1);
    try expect(my_list.items.len == 1);

    try my_list.append(1);
    try expect(my_list.items.len == 2);

    try my_list.append(1);
    try expect(my_list.items.len == 3);
}

test "add multiple items to array list" {
    var my_list = std.ArrayList(u8).init(allocator);
    defer my_list.deinit();

    const input_str = "this is the input string";

    for (input_str) |c| {
        try my_list.append(c);
    }

    try expect(my_list.items.len == input_str.len);
}
