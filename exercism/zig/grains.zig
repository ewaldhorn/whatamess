// const std = @import("std");

pub const ChessboardError = error{IndexOutOfBounds};

pub fn square(index: usize) ChessboardError!u64 {
    if (index < 1 or index > 64) {
        return error.IndexOutOfBounds;
    }
    var result: u64 = 1;

    for (1..index) |_| {
        result += result;
    }

    return result;
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
