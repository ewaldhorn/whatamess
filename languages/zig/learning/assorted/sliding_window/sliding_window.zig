const std = @import("std");
const print = std.debug.print;

// calculates the maximum value in the window size, across the array
// array is const since we won't be changing it
fn calculate_maximum(values: []const i32, window_size: isize) isize {
    var max: isize = -1;

    // if we don't have enough values to start with, what's the point even?
    // instant fail
    if (values.len < window_size) {
        return -1;
    }

    return max;
}

test "empty array" {
    const result = calculate_maximum(&[_]i32{}, 3);

    try std.testing.expect(result == -1);
}

test "too short array" {
    const result = calculate_maximum(&[_]i32{1}, 3);

    try std.testing.expect(result == -1);
}

test "too short array by one" {
    const result = calculate_maximum(&[_]i32{ 1, 2 }, 3);

    try std.testing.expect(result == -1);
}
