const std = @import("std");

pub fn mmap_file(path: []const u8) ![]align(std.mem.page_size) u8 {
    var file = try std.fs.createFileAbsolute(path, .{ .truncate = false, .read = true });
    defer file.close();
    const size = try file.getEndPos();

    var mapping = try std.os.mmap(null, size, std.os.PROT_READ | std.os.PROT_WRITE, std.os.MAP_SHARED, file.handle, 0);
    errdefer std.os.munmap(mapping);

    if (size < @sizeOf(u64)) {
        return error.FileTooSmall;
    }

    const data_end = size - @sizeOf(u64);
    const actual_hash = std.hash.CityHash64.hash(mapping[0..data_end]);
    const expected_hash = std.mem.readIntLittle(u64, mapping[data_end..]);
    
    if (actual_hash != expected_hash) {
        return error.FileSignatureInvalid;
    }

    return mapping;
}