const std = @import("std");
const expect = std.testing.expect;

// The standard library provides utilities for in-place sorting slices. Its basic usage is as follows.
test "sorting" {
    var data = [_]u8{ 10, 240, 0, 0, 10, 5 };
    std.mem.sort(u8, &data, {}, comptime std.sort.asc(u8));
    try expect(std.mem.eql(u8, &data, &[_]u8{ 0, 0, 5, 10, 10, 240 }));
    std.mem.sort(u8, &data, {}, comptime std.sort.desc(u8));
    try expect(std.mem.eql(u8, &data, &[_]u8{ 240, 10, 10, 5, 0, 0 }));
}

test "sorting floats" {
    var data = [_]f32{1.2, 2.4, 1.3, 1.1, 1.4};
    std.mem.sort(f32, &data, {}, comptime std.sort.desc(f32));

    try expect(std.mem.eql(f32, &data, &[_]f32{2.4, 1.4, 1.3, 1.2, 1.1}));
}

// std.sort.asc and .desc create a comparison function for the given type at comptime; 
// if non-numerical types should be sorted, the user must provide their own comparison function.
// std.sort.sort has a best case of O(n), and an average and worst case of O(n*log(n)).