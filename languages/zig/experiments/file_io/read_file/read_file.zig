const std = @import("std");

pub fn mmap_file(path: []const u8) ![]align(std.mem.page_size) u8 {
    var file = try std.fs.createFileAbsolute(path, .{ .truncate = false, .read = true });
    defer file.close();
    const size = try file.getEndPos();

    var mapping = try std.os.mmap(null, size, std.os.PROT.READ | std.os.PROT.WRITE, std.os.MAP.SHARED, file.handle, 0);
    errdefer std.os.munmap(mapping);

    if (size < @sizeOf(u64)) {
        return error.FileTooSmall;
    }

    const data_end = size - @sizeOf(u64);
    const actual_hash = std.hash.CityHash64.hash(mapping[0..data_end]);
    const expected_hash = std.mem.readInt(u64, mapping[data_end..], std.builtin.Endian.little);

    if (actual_hash != expected_hash) {
        return error.FileSignatureInvalid;
    }

    return mapping;
}

const expect = std.testing.expect;

test "file is too small" {
    const result = try mmap_file("./small.txt");
    try expect(std.mem.eql(u8, result, [_]u8{}));
}
