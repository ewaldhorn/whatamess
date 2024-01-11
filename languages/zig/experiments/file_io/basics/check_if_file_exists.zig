// Given the following list of files:
// .
// ├─ check_if_file_exists.zig
// └─ test.txt

const std = @import("std");

fn fileExists(path: []const u8) !bool {
    const file = std.fs.cwd().openFile(path, .{}) catch |err| switch (err) {
        error.FileNotFound => return false,
        else => return err,
    };

    file.close();

    return true;
}

test "Check if a file exists" {
    try std.testing.expect(try fileExists("test.txt"));
    try std.testing.expect(!try fileExists("does_not_exist.txt"));
}
