const std = @import("std");

fn splitStringIntoArrayOfCharacters(alloc : std.mem.Allocator, string: []const u8) ![]u8 {
    var list = try std.ArrayList(u8).initCapacity(alloc, string.len);
    defer list.deinit();

    for (string) |c| {
        list.appendAssumeCapacity(c);
    }

    return try list.toOwnedSlice();
}

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

test "can append a slice" {
    var my_list = std.ArrayList(u8).init(allocator);
    defer my_list.deinit();

    try my_list.appendSlice("This is a slice");
    try my_list.appendSlice(" ");
    try my_list.appendSlice("into a whole.");

    try expect(my_list.items.len == "This is a slice into a whole.".len);
}

test "remove and swap in an array list" {
    var my_list = std.ArrayList(i32).init(allocator);
    defer my_list.deinit();

    var n : i32 = 1;
     while (n <= 10) : (n += 1){
        try my_list.append(n);
    }

    try expect(my_list.items.len == 10);

    _ = my_list.orderedRemove(5);

    try expect(my_list.items.len == 9);
    try expect(my_list.items[0] == 1);
    try expect(my_list.items[8] == 10);

    const result = my_list.swapRemove(0);

    try expect(result == 1);
    try expect(my_list.items.len == 8);
    try expect(my_list.items[0] == 10);
}

test "pre-allocat memory for the array list" {
    var my_list = try std.ArrayList(u8).initCapacity(allocator, 12);
    defer my_list.deinit();

    try expect(my_list.capacity == 12);
    try expect(my_list.items.len == 0);

    for ("YOLO") |c| {
        my_list.appendAssumeCapacity(c);
    }

    try expect(my_list.capacity == 12);
    try expect(my_list.items.len == 4);
}

test "returning an array from an array list" {
    const source_string = "This is the input string";
    const result = try splitStringIntoArrayOfCharacters(allocator, source_string);
    defer allocator.free(result);

    try expect(result.len == source_string.len);
}