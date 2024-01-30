const std = @import("std");

pub fn buzz(val: u32, allocator: std.mem.Allocator) []const u8 {
    if (val % 15 == 0) {
        return std.fmt.allocPrint(allocator, "{s}", .{"FizzBuzz"}) catch "Error allocating memory.";
    } else if (val % 5 == 0) {
        return std.fmt.allocPrint(allocator, "{s}", .{"Buzz"}) catch "Error allocating memory.";
    } else if (val % 3 == 0) {
        return std.fmt.allocPrint(allocator, "{s}", .{"Fizz"}) catch "Error allocating memory.";
    }

    return std.fmt.allocPrint(allocator, "{d}", .{val}) catch "Error allocating memory.";    
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const test_allocator = std.testing.allocator;
const expect = std.testing.expect;
const eql = std.mem.eql;

test "get a number back - 1" {
    const result = buzz(1, test_allocator);
    defer test_allocator.free(result);

    try expect(eql(u8, "1", result));
}

test "get FizzBuzz back" {
    const result = buzz(15, test_allocator);
    defer test_allocator.free(result);

    try expect(eql(u8, "FizzBuzz", result));
}

test "get Fizz back" {
    const result = buzz(3, test_allocator);
    defer test_allocator.free(result);

    try expect(eql(u8, "Fizz", result));
}

test "get Buzz back" {
    const result = buzz(5, test_allocator);
    defer test_allocator.free(result);

    try expect(eql(u8, "Buzz", result));
}

