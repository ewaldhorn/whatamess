const std = @import("std");

test "create and write to a file, then append." {
    var file = try std.fs.cwd().createFile("test_append.txt", .{});

    const wrote = try file.write("A bunch of bytes that makes sense to human readers.\n");

    std.debug.print("We ended up writing {d} bytes to the file just now.\n", .{wrote});

    const Writer = file.writer();

    const phone_number = 911;

    try Writer.print("Help! Someone, please call {d}!\n\n", .{phone_number});
    file.close();

    file = try std.fs.cwd().createFile("test_append.txt", .{ .truncate = false });
    defer file.close();

    try file.writer().print("This is fun!\n", .{});
}
