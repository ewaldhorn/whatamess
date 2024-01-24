const std = @import("std");

pub fn main() !void {
    const birthYear: u32 = 1978;
    const currentYear: u32 = 2023;
    var counter: u8 = 0;

    std.debug.print("The current state:\n", .{});
    std.debug.print("birthYear: {}\ncurrentYear: {}\ncounter: {}\n", .{ birthYear, currentYear, counter });

    while (counter < (currentYear - birthYear)) {
        counter += 1;
    }

    std.debug.print("\nThe new state:\n", .{});
    std.debug.print("birthYear: {}\ncurrentYear: {}\ncounter: {}\n", .{ birthYear, currentYear, counter });
}
