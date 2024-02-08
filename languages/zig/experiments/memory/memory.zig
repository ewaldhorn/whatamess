const std = @import("std");

fn createAnArrayList(allocator: std.mem.Allocator) !std.ArrayList(i32) {
    var list = std.ArrayList(i32).init(allocator);
    errdefer list.deinit();

    try list.append(123);

    return list;
}

fn modifyExistingArrayList(list: *std.ArrayList(i32)) !void {
    try list.append(10);
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS

test "can receive an array list and modify it" {
    var list = try createAnArrayList(std.testing.allocator);
    defer list.deinit();

    try std.testing.expect(list.items.len == 1);
    try std.testing.expect(list.items[0] == 123);

    try list.append(456);
    try std.testing.expect(list.items.len == 2);

    const value = list.items[1];
    try std.testing.expect(value == 456);
}

test "can modify existing array list" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit();

    try modifyExistingArrayList(&list);
    try std.testing.expect(list.items.len == 1);
    try std.testing.expect(list.items[0] == 10);

    try list.append(20);
    try std.testing.expect(list.items.len == 2);
    try std.testing.expect(list.items[1] == 20);
    try std.testing.expect(list.getLast() == 20);
}