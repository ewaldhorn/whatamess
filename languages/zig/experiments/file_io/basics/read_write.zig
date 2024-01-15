const std = @import("std");

test "Read/write file" {
    const file = try std.fs.cwd().openFile("read_write.txt", .{ .mode = std.fs.File.OpenMode.read_write });
    defer file.close();

    const file_size = try file.getEndPos();

    while (true) {
        const current_position = try file.getPos();

        if (current_position + 4 >= file_size)
            break;

        try file.seekBy(4);

        try file.writer().writeByte('a');
    }
}
