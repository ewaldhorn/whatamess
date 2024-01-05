const testing = @import("std").testing;

// function to make it easier to write tests - less of the "@as" boilerplate needed
pub fn expectEqual(expected: anytype, actual: anytype) void {
    testing.expectEqual(@as(@TypeOf(actual), expected), actual);
}

test "expectEqual correctly coerces types that std.testing.expectEqual does not" {
    const int_value: u8 = 2;
    expectEqual(2, int_value);

    const optional_value: ?u8 = null;
    expectEqual(null, optional_value);

    const Enum = enum { One, Two };
    const enum_value = Enum.One;
    expectEqual(.One, enum_value);
}