// Given the following list of files:
// .
// ├─ file.zig
// └─ sample.txt

const std = @import("std");

test "Delete file" {
    var file = try std.fs.cwd().createFile("sample.txt", .{});
    _ = try file.write("A bunch of bytes that makes sense to human readers.\n");
    file.close();

    std.fs.cwd().deleteFile("sample.txt") catch |err| switch (err) {
        error.FileNotFound => {},
        else => return err,
    };

    const has_been_deleted = std.fs.cwd().openFile("sample.txt", .{});

    try std.testing.expectError(error.FileNotFound, has_been_deleted);
}
