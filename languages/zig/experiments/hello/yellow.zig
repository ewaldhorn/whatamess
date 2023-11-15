const print = @import("std").debug.print;

pub fn main() void {
    print("Yellow there {s}\n", .{"Zig!"});
}
