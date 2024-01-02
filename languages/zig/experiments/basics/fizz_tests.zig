const std = @import("std");
const expect = @import("std").testing.expect;

fn fizzit(num : i32) []const u8 {

            if (num % 3 == 0 and num % 5 == 0) {
            return "FizzBuzz";
        } else if (num % 3 == 0) {
            return "Fizz";
        } else if (num % 5 == 0) {
            return "Buzz";
        } else {
            return "{d}\n", .{i};
        }

    return "lol";
}

test "1 returns 1" {
    try expect(std.mem.eql(u8, fizzit(1), "lol"));
}