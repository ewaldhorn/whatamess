const std = @import("std");
const embedded_content = @embedFile("./embed_me.txt");

pub fn readInputFile(allocator: std.mem.Allocator) ![]const u8 {
    const file = try std.fs.cwd().openFile("./embed_me.txt", .{});
    defer file.close();
    const stat = try file.stat();
    const fileSize = stat.size;
    return try file.reader().readAllAlloc(allocator, fileSize);
}

pub fn splitStringIntoLines(lines : *std.ArrayList([]const u8), content: []const u8) !void {
    var readIter = std.mem.tokenize(u8, content, "\n");

    while (readIter.next()) |line| {
        try lines.append(line);
    }
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS

test "can read embedded content" {
    const result = embedded_content.len;

    try std.testing.expect(result == 55);
}

test "can read external file content" {
    const result = try readInputFile(std.testing.allocator);
    defer std.testing.allocator.free(result);

    try std.testing.expect(result.len == 55);
}

test "can split embedded content" {
    var lines = std.ArrayList([]const u8).init(std.testing.allocator);
    defer lines.deinit();

    var readIter = std.mem.tokenize(u8, embedded_content, "\n");
    while (readIter.next()) |line| {
        try lines.append(line);
    }

    try std.testing.expect(lines.items.len == 6);
}

test "can read and split external file content" {
    const result = try readInputFile(std.testing.allocator);
    defer std.testing.allocator.free(result);

    var lines = std.ArrayList([]const u8).init(std.testing.allocator);
    defer lines.deinit();

    var readIter = std.mem.tokenize(u8, result, "\n");
    while (readIter.next()) |line| {
        try lines.append(line);
    }

    try std.testing.expect(lines.items.len == 6);
}

test "can split embedded content using function" {
    var lines = std.ArrayList([]const u8).init(std.testing.allocator);
    defer lines.deinit();

    try splitStringIntoLines(&lines, embedded_content);

    try std.testing.expect(lines.items.len == 6);
}
