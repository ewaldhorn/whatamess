const std = @import("std");
pub fn recite(allocator: std.mem.Allocator, words: []const []const u8) (std.fmt.AllocPrintError || std.mem.Allocator.Error)![][]u8 {
    var output = std.ArrayList([]u8).init(allocator);

    if (words.len > 1) {
        for (0..words.len - 1) |idx| {
            try output.append(try std.fmt.allocPrint(allocator, "For want of a {s} the {s} was lost.\n", .{ words[idx], words[idx + 1] }));
        }
    }

    if (words.len > 0) {
        try output.append(try std.fmt.allocPrint(allocator, "And all for the want of a {s}.\n", .{words[0]}));
    }

    return output.toOwnedSlice();
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const testing = std.testing;

fn free(slices: [][]u8) void {
    for (slices) |s| {
        testing.allocator.free(s);
    }
    testing.allocator.free(slices); // No problem when `slices` has zero length.
}

test "zero pieces" {
    const input = [_][]const u8{};
    const expected = [_][]const u8{};
    const actual = try recite(testing.allocator, &input);
    defer free(actual);
    try testing.expectEqualSlices([]const u8, &expected, actual);
}

test "one piece" {
    const input = [_][]const u8{
        "nail",
    };
    const expected = [_][]const u8{
        "And all for the want of a nail.\n",
    };
    const actual = try recite(testing.allocator, &input);
    defer free(actual);
    for (expected, 0..) |expected_slice, i| {
        try testing.expectEqualSlices(u8, expected_slice, actual[i]);
    }
}

test "two pieces" {
    const input = [_][]const u8{
        "nail",
        "shoe",
    };
    const expected = [_][]const u8{
        "For want of a nail the shoe was lost.\n",
        "And all for the want of a nail.\n",
    };
    const actual = try recite(testing.allocator, &input);
    defer free(actual);
    for (expected, 0..) |expected_slice, i| {
        try testing.expectEqualSlices(u8, expected_slice, actual[i]);
    }
}

test "three pieces" {
    const input = [_][]const u8{
        "nail",
        "shoe",
        "horse",
    };
    const expected = [_][]const u8{
        "For want of a nail the shoe was lost.\n",
        "For want of a shoe the horse was lost.\n",
        "And all for the want of a nail.\n",
    };
    const actual = try recite(testing.allocator, &input);
    defer free(actual);
    for (expected, 0..) |expected_slice, i| {
        try testing.expectEqualSlices(u8, expected_slice, actual[i]);
    }
}

test "full proverb" {
    const input = [_][]const u8{
        "nail",
        "shoe",
        "horse",
        "rider",
        "message",
        "battle",
        "kingdom",
    };
    const expected = [_][]const u8{
        "For want of a nail the shoe was lost.\n",
        "For want of a shoe the horse was lost.\n",
        "For want of a horse the rider was lost.\n",
        "For want of a rider the message was lost.\n",
        "For want of a message the battle was lost.\n",
        "For want of a battle the kingdom was lost.\n",
        "And all for the want of a nail.\n",
    };
    const actual = try recite(testing.allocator, &input);
    defer free(actual);
    for (expected, 0..) |expected_slice, i| {
        try testing.expectEqualSlices(u8, expected_slice, actual[i]);
    }
}

test "four pieces modernized" {
    const input = [_][]const u8{
        "pin",
        "gun",
        "soldier",
        "battle",
    };
    const expected = [_][]const u8{
        "For want of a pin the gun was lost.\n",
        "For want of a gun the soldier was lost.\n",
        "For want of a soldier the battle was lost.\n",
        "And all for the want of a pin.\n",
    };
    const actual = try recite(testing.allocator, &input);
    defer free(actual);
    for (expected, 0..) |expected_slice, i| {
        try testing.expectEqualSlices(u8, expected_slice, actual[i]);
    }
}
