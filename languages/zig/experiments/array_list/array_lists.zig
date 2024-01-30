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

test "simulate a stack" {
    var my_stack = std.ArrayList(u8).init(allocator);
    defer my_stack.deinit();

    try my_stack.append(2);
    try my_stack.append(4);
    try my_stack.append(6);
    try my_stack.append(8);

    try expect(my_stack.items.len == 4);

    _ = my_stack.pop();

    try expect(my_stack.items.len == 3);
    try expect(my_stack.items[0] == 2);
    try expect(my_stack.items[1] == 4);
    try expect(my_stack.items[2] == 6);
}

test "use it like a writer if it's of the u8 type" {
    var my_list = std.ArrayList(u8).init(allocator);
    defer my_list.deinit();

    const writer = my_list.writer();

    _ = try writer.print("A formatted string with {d} as well.", .{9});

    try expect(my_list.items.len == 34);
}

test "can do stuff in reverse" {
    var my_list = std.ArrayList(u8).init(allocator);
    defer my_list.deinit();

    for ("this is a string") |c| {
        try my_list.append(c);
    }

    var my_second_list = std.ArrayList(u8).init(allocator);
    defer my_second_list.deinit();

    while (my_list.popOrNull()) |c| {
        try my_second_list.append(c);
    }

    try expect(my_list.items.len == 0);
    try expect(my_second_list.items[my_second_list.items.len-1] == 't');
    try expect(my_second_list.items[my_second_list.items.len-2] == 'h');
    try expect(my_second_list.items[my_second_list.items.len-3] == 'i');
    try expect(my_second_list.items[my_second_list.items.len-4] == 's');
    try expect(my_second_list.items[my_second_list.items.len-5] == ' ');
    try expect(my_second_list.items[my_second_list.items.len-6] == 'i');
}