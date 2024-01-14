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

test "brute empty array" {
    const result = calculate_maximum_brute(&[_]i32{}, 3);

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
