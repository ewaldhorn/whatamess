const std = @import("std");
const expect = std.testing.expect;

const Status = enum {
    starting,
    started,
    running,
    stopping,
    stopped,

    fn isRunning(s: Status) bool {
        return s == .running;
    }
};

pub fn main() void {
    const myVal = Status.running;

    std.debug.print("{s}.\n", .{@tagName(myVal)}); // get the name of the enum as a string
}

// ========================================================================================== TESTS
test "basic enum test" {
    const val = Status.started;

    try expect(val.isRunning() == false);
}

test "enum variable test" {
    var val = Status.starting;

    try expect(val == Status.starting);

    val = Status.started;
    try expect(val.isRunning() == false);

    val = Status.running;
    try expect(val.isRunning() == true);
}

test "enum tag name" {
    var val = Status.starting;

    try expect(std.mem.eql(u8, "starting", @tagName(val)));

    val = Status.running;
    try expect(std.mem.eql(u8, "running", @tagName(val)));
}