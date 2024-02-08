const std = @import("std");

pub const ChessboardError = error{IndexOutOfBounds};

pub fn square(index: usize) ChessboardError!u64 {
    if (index < 1 or index > 64) {
        return error.IndexOutOfBounds;
    }

    return std.math.pow(u64, 2, index - 1);
}

pub fn total() u64 {
    var result: u64 = 0;

    var idx: usize = 1;
    while (idx <= 64) : (idx += 1) {
        if (square(idx)) |n| {
            // std.debug.print("idx {} and {}\n", .{ idx, n });
            result += n;
        } else |_| {}
    }

    return result;
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const expect = std.testing.expect;

test "square of 0 errors" {
    const result = square(0);

    try expect(result == ChessboardError.IndexOutOfBounds);
}


test "square of 65 errors" {
    const result = square(65);

    try expect(result == ChessboardError.IndexOutOfBounds);
}