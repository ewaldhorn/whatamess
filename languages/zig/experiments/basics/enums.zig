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