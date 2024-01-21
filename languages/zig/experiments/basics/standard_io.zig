const std = @import("std");
const expect = std.testing.expect;
const eql = std.mem.eql;

const ArrayList = std.ArrayList;
const test_allocator = std.testing.allocator;

// std.io.Writer and std.io.Reader provide standard ways of making use of IO. std.ArrayList(u8) has 
// a writer method which gives us a writer. Let's use it.
test "io writer usage" {
    var list = ArrayList(u8).init(test_allocator);
    defer list.deinit();

    const bytes_written = try list.writer().write(
        "Hello World!",
    );

    try expect(bytes_written == 12);
    try expect(eql(u8, list.items, "Hello World!"));
}

// Here we will use a reader to copy the file's contents into an allocated buffer. The second 
// argument of readAllAlloc is the maximum size that it may allocate; if the file is larger than 
// this, it will return error.StreamTooLong.
test "io reader usage" {
    const message = "Hello File!";

    const file = try std.fs.cwd().createFile(
        "junk_file2.txt",
        .{ .read = true },
    );
    defer file.close();

    try file.writeAll(message);
    try file.seekTo(0);

    const contents = try file.reader().readAllAlloc(
        test_allocator,
        message.len,
    );
    defer test_allocator.free(contents);

    try expect(eql(u8, contents, message));
}