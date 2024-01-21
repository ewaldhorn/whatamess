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

// A common usecase for readers is to read until the next line (e.g. for user input). Here we will 
// do this with the std.io.getStdIn() file.
fn nextLine(reader: anytype, buffer: []u8) !?[]const u8 {
    const line = (try reader.readUntilDelimiterOrEof(
        buffer,
        '\n',
    )) orelse return null;

    // trim annoying windows-only carriage return character
    if (@import("builtin").os.tag == .windows) {
        return std.mem.trimRight(u8, line, "\r");
    } else {
        return line;
    }
}

test "read until next line" {
    const stdout = std.io.getStdOut();
    const stdin = std.io.getStdIn();

    try stdout.writeAll("\n\nEnter your name: ");

    var buffer: [100]u8 = undefined;
    const input = (try nextLine(stdin.reader(), &buffer)).?;
    try stdout.writer().print(
        "Your name is: \"{s}\".\n\n",
        .{input},
    );
}

// An std.io.Writer type consists of a context type, error set, and a write function. The write 
// function must take in the context type and a byte slice. The write function must also return an 
// error union of the Writer type's error set and the number of bytes written. Let's create a type 
// that implements a writer.
// ------------------------------------------------------------------------------------------------
// Don't create a type like this! Use an arraylist with a fixed buffer allocator
const MyByteList = struct {
    data: [100]u8 = undefined,
    items: []u8 = &[_]u8{},

    const Writer = std.io.Writer(
        *MyByteList,
        error{EndOfBuffer},
        appendWrite,
    );

    fn appendWrite(
        self: *MyByteList,
        data: []const u8,
    ) error{EndOfBuffer}!usize {
        if (self.items.len + data.len > self.data.len) {
            return error.EndOfBuffer;
        }
        std.mem.copyForwards(u8,self.data[self.items.len..],data);
        self.items = self.data[0 .. self.items.len + data.len];
        return data.len;
    }

    fn writer(self: *MyByteList) Writer {
        return .{ .context = self };
    }
};

test "custom writer" {
    var bytes = MyByteList{};
    _ = try bytes.writer().write("Hello");
    _ = try bytes.writer().write(" Writer!");
    try expect(eql(u8, bytes.items, "Hello Writer!"));
}