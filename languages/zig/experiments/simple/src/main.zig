const std = @import("std");

pub fn main() !void {
    var birthYear: u32 = 1978;
    var currentYear: u32 = 2023;
    var counter: u8 = 0;

    std.debug.print("The current state:\n", .{});
    std.debug.print("birthYear: {}\ncurrentYear: {}\ncounter: {}\n", .{ birthYear, currentYear, counter });

    while (counter < (currentYear - birthYear)) {
        counter += 1;
    }

    std.debug.print("\nThe new state:\n", .{});
    std.debug.print("birthYear: {}\ncurrentYear: {}\ncounter: {}\n", .{ birthYear, currentYear, counter });
}
