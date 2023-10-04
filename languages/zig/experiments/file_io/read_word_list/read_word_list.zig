const std = @import("std");

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    std.debug.print("The words are:\n", .{});

    if (loadWordList(allocator)) |words| {
        std.debug.print("Words contains {} entries.\n", .{words.len});
    } else |err| {
        std.debug.print("Error: {}\n", .{err});
    }
}

fn showWorkingDirectory() void {
    // allocate a large enough buffer to store the cwd
    var buf: [std.fs.MAX_PATH_BYTES]u8 = undefined;

    // getcwd writes the path of the cwd into buf and returns a slice of buf with the len of cwd
    const cwd = try std.os.getcwd(&buf);

    // print out the cwd
    std.debug.print("cwd: {s}", .{cwd});
}

fn loadWordList(allocator: std.mem.Allocator) ![][]u8 {
    var array = std.ArrayList([]u8).init(allocator);

    var file = try std.fs.cwd().openFile("words.txt", .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();
    var buf: [1024]u8 = undefined;

    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        std.debug.print("{s}\n", .{line});
        try array.append(line);
    }

    return array.toOwnedSlice();
}

const testing = std.testing;
test "try to load word list" {
    if (loadWordList(testing.allocator)) |words| {
        std.debug.print("Words contains {} entries.\n", .{words.len});
    } else |err| {
        std.debug.print("An error occurred : {}", .{err});
        _ = err catch {};
    }
}
