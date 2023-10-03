const std = @import("std");

pub fn main() !void {
    // Integers
    var userCount: i32 = 10;
    _ = try reportSize("userCount", @TypeOf(userCount));

    // Floating-points
    const userAverage: f64 = 0.001;
    _ = try reportSize("userAverage", @TypeOf(userAverage));

    // Booleans
    const isSelected = false;
    _ = try reportSize("isSelected", @TypeOf(isSelected));

    // Structs
    const User = struct { age: i8 };

    _ = try reportSize("User", @TypeOf(User));
}

pub fn reportSize(name: []const u8, comptime input: type) !void {
    std.debug.print("{s} is: {}, Size of {s} is {} in byte(s).\n", .{ name, input, name, @sizeOf(input) });
}
