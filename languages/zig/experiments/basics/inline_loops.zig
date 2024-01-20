const std = @import("std");
const assert = std.debug.assert;

const Data = struct {
    foo: Foo,
    bytes: [8]u8,
    ok: bool,
};

const Foo = enum {hello, world};

pub fn main() void {
    const d: Data = .{
        .foo = .world,
        .bytes = "abcdefgh".*,
        .ok = true,
    };
    dump(d);
}

fn dump(data: anytype) void {
    const T = @TypeOf(data);
    inline for (@typeInfo(T).Struct.fields) |field| {
        std.debug.print("{any}\n", .{@field(data, field.name)});
    }
}

test "inline for" {
    const types = [_]type{ i32, f32, u8, bool };
    var sum: usize = 0;
    inline for (types) |T| sum += @sizeOf(T);
    try std.testing.expect(sum == 10);
}