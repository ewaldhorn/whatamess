const std = @import("std");
const expect = std.testing.expect;

test "Variadic function" {
    const twelve: u8 = 12;
    const fifty_one: u8 = 51;

    const foo = sum(u8, .{ 4, 1, twelve, fifty_one, 33, 41, 99 });
    try expect(foo == 241);

    const asdf: f32 = 219.231;

    const bar = sum(f32, .{ 21.41, 132.111, 521.122224, asdf, 10 });
    try expect(bar < 1000);
}

fn sum(comptime T: type, args: anytype) T {
    var answer: T = 0;

    // The trick here is to iterate over the `args` tuple at compile-time.
    comptime {
        var i: usize = 0;

        while (i < args.len) : (i += 1) {
            answer += args[i];
        }
    }

    return answer;
}
