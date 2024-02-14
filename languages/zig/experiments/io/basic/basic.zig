const std = @import("std");

pub fn readInputFile(allocator: std.mem.Allocator) ![]const u8 {
    const file = try std.fs.cwd().openFile("text.txt", .{});
    defer file.close();
    
    const stat = try file.stat();
    const fileSize = stat.size;

    return try file.reader().readAllAlloc(allocator, fileSize);
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const test_allocator = std.testing.allocator;

test "can read the file" {
    const result = try readInputFile(test_allocator);
    defer test_allocator.free(result);
    
    try std.testing.expectEqual(161, result.len);
}