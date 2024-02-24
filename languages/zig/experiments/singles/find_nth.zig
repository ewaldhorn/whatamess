const std = @import("std");

// Given an array of integers, find the Nth largest one.
// For example:
// Input 4,2,9,7,5,6,7,1,3
// N = 4
// Answer = 6
// How? 1st = 9, 2nd = 7, 3rd = 7, 4th = 6
fn findNthLargest(source: []const u8, target: u8) u8 {
    if (target > source.len) {
        return 0;
    }

    var buffer : [1000]u8 = undefined;
    const slice : []u8 = buffer[0..source.len];
    std.mem.copyForwards(u8, slice, source);
    std.mem.sort(u8, slice, {}, comptime std.sort.asc(u8));

    return slice[slice.len - target];
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const expectEqual = std.testing.expectEqual;

test "invalid test case" {
    const input = [_]u8{4,2,9,7,5,6};
    const target = 7;

    try expectEqual(0, findNthLargest(&input, target));
}

test "simple test case" {
    const input = [_]u8{4,2,9,7,5,6,7,1,3};
    const target = 4;

    try expectEqual(6, findNthLargest(&input, target));
}

test "unlikely test case" {
    const input = [_]u8{3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3};
    const target = 11;

    try expectEqual(3, findNthLargest(&input, target));
}
