const std = @import("std");

const OpenError = error {
	AccessDenied,
	NotFound,
};

fn returnError(num: usize) OpenError!void {
    if (num < 1000) {
        return OpenError.AccessDenied;
    }
}

test "test that no error is returned" {
    try std.testing.expect(returnError(10_000) != OpenError.AccessDenied);
}

test "that an error is returned" {
    try std.testing.expect(returnError(99) == OpenError.AccessDenied);
}