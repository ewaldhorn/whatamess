const std = @import("std");

const InputError = error {
    TooShort,TooLong
};

const ValidationError = error {
    NotFound, NotActive
};

// combine all our errors
const AllErrors = InputError || ValidationError;

fn readInput(input: usize) !usize {
    switch (input) {
        0 => return 0,
        1 => return AllErrors.TooShort,
        2 => return AllErrors.NotFound,
        else => return AllErrors.NotFound
    }
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const expectEqual = std.testing.expectEqual;

test "0" {
    const result = try readInput(0);

    try expectEqual(0, result);
}

test "catch error" {
    const result = readInput(1) catch 99;

    try expectEqual(99, result);
}

test "catch error alternative" {
    const result: AllErrors!usize = readInput(2) catch |err| {
        std.debug.print("an error occured: {}", .{err});
        return AllErrors.NotActive;
    };

    try std.testing.expectError(AllErrors.NotActive, result);
}

// test "1" {
//     const result = try readInput(1);
//     std.debug.print("\n\n>>>>>>>>>{d}\n\n", .{result});
//     try expectEqual(result, AllErrors.TooShort);
// }