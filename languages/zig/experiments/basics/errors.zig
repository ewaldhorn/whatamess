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

fn doesntReturnError(num: usize) bool {
    var result = true;

    _ = returnError(num) catch |err| {
        std.debug.print("\n>>>>>>\nWe have {}\n", .{err});
        // if (err == OpenError.AccessDenied) {
        //     result = false;
        // }

        result = false;
    };

    return result;
}

test "test that no error is returned" {
    try std.testing.expect(returnError(10_000) != OpenError.AccessDenied);
}

test "that an error is returned" {
    try std.testing.expect(returnError(99) == OpenError.AccessDenied);
}

test "error trap returns true because it worked" {
    try std.testing.expect(doesntReturnError(900) == true);
}

test "error trap returns false because it failed" {
    try std.testing.expect(doesntReturnError(9_000) == false);
}