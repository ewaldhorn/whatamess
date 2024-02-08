const std = @import("std");

const State = enum {
    Pending,
    Processing,
    Packaging,
    Shipped,

    pub fn hasShipped(self: State) bool {
        return self == .Shipped;
    }
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

test "enums have int values" {
    const status = State.Packaging;
    const value = @intFromEnum(status);

    try expect(value == 2);
}

test "can use int value to get to enum value" {
    const value : State = @enumFromInt(3);

    try expect(value == State.Shipped);
}

test "can call bound functions on enum" {
    const status = State.Packaging;

    try expect(status.hasShipped() == false);
    try expect(State.hasShipped(status) == false);
}