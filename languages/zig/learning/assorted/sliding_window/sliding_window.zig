const std = @import("std");
const print = std.debug.print;

// calculates the maximum value in the window size, across the array
// array is const since we won't be changing it
// this is a brute-force method
fn calculate_maximum_brute(values: []const i32, window_size: usize) isize {
    var max: isize = -1;

    // if we don't have enough values to start with, what's the point even?
    // instant fail
    if (values.len < window_size) {
        return -1;
    }

    const position_cap: usize = values.len - window_size;
    var position: usize = 0;

    while (position < position_cap) : (position += 1) {
        var tmp = values[position];

        var tmp_position: usize = 1;

        while (tmp_position < window_size) : (tmp_position += 1) {
            tmp += values[position + tmp_position];
        }

        if (tmp > max) {
            max = tmp;
        }
    }

    return max;
}

// uses the sliding-window principle to calculate the maximum value in the specified window size
// uses less loops than the brute-force version
// as the window moves right, the FIRST value of the previous window is subtracted and the NEW
// end value of the new window is added. This creates a new total for the current window.
fn calculate_maximum_sliding(values: []const i32, window_size: usize) isize {
    var max: isize = -1;

    if (values.len < window_size) {
        return -1;
    }

    var start: usize = 0;
    var end: usize = 0;
    var window_max: isize = 0;

    // get the starting values
    while (end < window_size) {
        window_max += values[end];
        end += 1;
    }
    max = window_max;

    while (start < values.len - window_size) {
        window_max = window_max - values[start] + values[end];

        if (window_max > max) {
            max = window_max;
        }

        start += 1;
        end += 1;
    }

    return max;
}

test "brute empty array" {
    const result = calculate_maximum_brute(&[_]i32{}, 3);

    try std.testing.expect(result == -1);
}

test "sliding empty array" {
    const result = calculate_maximum_sliding(&[_]i32{}, 3);

    try std.testing.expect(result == -1);
}

test "brute too short array" {
    const result = calculate_maximum_brute(&[_]i32{1}, 3);

    try std.testing.expect(result == -1);
}

test "brute too short array by one" {
    const result = calculate_maximum_brute(&[_]i32{ 1, 2 }, 3);

    try std.testing.expect(result == -1);
}

test "sliding too short array" {
    const result = calculate_maximum_sliding(&[_]i32{1}, 3);

    try std.testing.expect(result == -1);
}

test "sliding too short array by one" {
    const result = calculate_maximum_sliding(&[_]i32{ 1, 2 }, 3);

    try std.testing.expect(result == -1);
}

test "brute output should be 12" {
    const result = calculate_maximum_brute(&[_]i32{ 1, 2, 6, 2, 4, 1 }, 3);

    try std.testing.expect(result == 12);
}

test "brute output should be 8" {
    const result = calculate_maximum_brute(&[_]i32{ 1, 2, 6, 2, 4, 1 }, 2);

    try std.testing.expect(result == 8);
}

test "brute output should be 27" {
    const result = calculate_maximum_brute(&[_]i32{ 1, 2, 6, 2, 4, 1, 4, 5, 3, 9, 8, 5, 7, 9, 9, 9, 3, 4, 8, 8, 8, 2, 1, 4, 5 }, 3);

    try std.testing.expect(result == 27);
}

test "brute output should be 18" {
    const result = calculate_maximum_brute(&[_]i32{ 1, 2, 6, 2, 4, 1, 4, 5, 3, 9, 8, 5, 7, 9, 9, 9, 3, 4, 8, 8, 8, 2, 1, 4, 5 }, 2);

    try std.testing.expect(result == 18);
}

test "sliding output should be 12" {
    const result = calculate_maximum_sliding(&[_]i32{ 1, 2, 6, 2, 4, 1 }, 3);

    try std.testing.expect(result == 12);
}

test "sliding output should be 8" {
    const result = calculate_maximum_sliding(&[_]i32{ 1, 2, 6, 2, 4, 1 }, 2);

    try std.testing.expect(result == 8);
}

test "sliding output should be 27" {
    const result = calculate_maximum_sliding(&[_]i32{ 1, 2, 6, 2, 4, 1, 4, 5, 3, 9, 8, 5, 7, 9, 9, 9, 3, 4, 8, 8, 8, 2, 1, 4, 5 }, 3);

    try std.testing.expect(result == 27);
}

test "sliding output should be 18" {
    const result = calculate_maximum_sliding(&[_]i32{ 1, 2, 6, 2, 4, 1, 4, 5, 3, 9, 8, 5, 7, 9, 9, 9, 3, 4, 8, 8, 8, 2, 1, 4, 5 }, 2);

    try std.testing.expect(result == 18);
}
