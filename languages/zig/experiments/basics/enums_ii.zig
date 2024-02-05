const std = @import("std");

const EnumType = enum(u8){
    EnumOne,
    EnumTwo,
    EnumThree = 5
};

pub fn main() void {
    std.debug.print("One: {}\n", .{EnumType.EnumOne});
    std.debug.print("Two?: {}\n", .{EnumType.EnumTwo == .EnumTwo});
    std.debug.print("Three?: {}\n", .{@intFromEnum(EnumType.EnumThree) == 5});
}
