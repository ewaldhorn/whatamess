const std = @import("std");

pub fn main() !void {
    const magicModuleHeader = [_]u8{ 0x00, 0x61, 0x73, 0x6d };
    const moduleVersion = [_]u8{ 0x01, 0x00, 0x00, 0x00 };

    var file = try std.fs.cwd().createFile("noop.wasm", .{});
    defer file.close();

    _ = try file.write(&magicModuleHeader);
    _ = try file.write(&moduleVersion);
}
