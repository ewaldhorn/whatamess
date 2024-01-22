const std = @import("std");
const expect = std.testing.expect;

fn add_upp(values : []const u32) u64 {
    var total : u64 = 0;

    for (values) |v| {
        total += v;
    }

    return total;
}

test "empty array" {
    const answer = add_upp(&[_]u32{});

    try expect(answer == 0);
}

test "array {0}" {
    const answer = add_upp(&[_]u32{0});

    try expect(answer == 0);
}

test "array {123}" {
    const answer = add_upp(&[_]u32{123});

    try expect(answer == 123);
}

test "array {1,2}" {
    const answer = add_upp(&[_]u32{1,2});

    try expect(answer == 3);
}

test "array {1,1,1,1,1}" {
    const answer = add_upp(&[_]u32{1,1,1,1,1});

    try expect(answer == 5);
}

test "array {1000000001,1000000002,1000000003,1000000004,1000000005}" {
    const answer = add_upp(&[_]u32{1000000001,1000000002,1000000003,1000000004,1000000005});

    try expect(answer == 5000000015);
}