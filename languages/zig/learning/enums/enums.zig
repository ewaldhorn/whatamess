const std = @import("std");

const State = enum {
    Pending,
    Processing,
    Packaging,
    Shipped
};

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const expect = std.testing.expect;

test "can create an enum" {
    const status = State.Pending;
    const status_alternate : State = .Processing;

    try expect(status == State.Pending);
    try expect(status_alternate == State.Processing);
}