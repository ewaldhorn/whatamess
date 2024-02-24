const std = @import("std");

// Given an array of integers and a target, return the index of the first and last occurences
// of the target number in the array.  The target could occur 0+ times in the source array.
// If not found, return -1 for that position.
// Expected result = [first, last]
fn findFirstLast(source: []const u8, target: u8) [2]isize {
    var first: isize = -1;
    var last: isize = -1;

    for (0..source.len) |idx| {
        if (source[idx] == target) {
            if (first == -1) {
                first = @as(isize, @intCast(idx));
            } else {
                last = @as(isize, @intCast(idx));
            }
        }
    }

    return [_]isize{ first, last };
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const expectEqual = std.testing.expectEqual;

const all_empty = [_]isize{ -1, -1 };

test "empty array" {
    try expectEqual(all_empty, findFirstLast(&[_]u8{}, 0));
}

test "valid array, no point (1)" {
    try expectEqual(all_empty, findFirstLast(&[_]u8{ 2, 3, 4, 5, 6, 7, 8 }, 1));
}

test "valid array, one point (2)" {
    try expectEqual([_]isize{ 0, -1 }, findFirstLast(&[_]u8{ 2, 4, 5, 5, 5, 5, 5, 7, 9, 9 }, 2));
}

test "valid array, one point (4)" {
    try expectEqual([_]isize{ 1, -1 }, findFirstLast(&[_]u8{ 2, 4, 5, 5, 5, 5, 5, 7, 9, 9 }, 4));
}

test "valid array, one point (7)" {
    try expectEqual([_]isize{ 7, -1 }, findFirstLast(&[_]u8{ 2, 4, 5, 5, 5, 5, 5, 7, 9, 9 }, 7));
}

test "valid array, two points (5)" {
    try expectEqual([_]isize{ 2, 6 }, findFirstLast(&[_]u8{ 2, 4, 5, 5, 5, 5, 5, 7, 9, 9 }, 5));
}

test "valid array, two points (9)" {
    try expectEqual([_]isize{ 8, 9 }, findFirstLast(&[_]u8{ 2, 4, 5, 5, 5, 5, 5, 7, 9, 9 }, 9));
}

test "valid array, two points (4)" {
    try expectEqual([_]isize{ 1, 9 }, findFirstLast(&[_]u8{ 2, 4, 5, 5, 5, 5, 5, 7, 9, 4 }, 4));
}
