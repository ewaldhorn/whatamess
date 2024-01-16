const std = @import("std");
const expect = std.testing.expect;

test "Variadic function" {
    var twelve: u8 = 12;
    var fifty_one: u8 = 51;

    const foo = sum(u8, .{ 4, 1, twelve, fifty_one, 33, 41, 99 });
    try expect(foo == 241);

    var asdf: f32 = 219.231;

    const bar = sum(f32, .{ 21.41, 132.111, 521.122224, asdf, 10 });
    try expect(bar < 1000);
}

fn sum(comptime T: type, args: anytype) T {
    const type_info = @typeInfo(T);
    switch (type_info) {
        .Int,
        .Float => {},
        else => @compileError("Expected integer or floating type, found " ++ @typeName(T)),
    }

    const ArgsType = @TypeOf(args);
    if (@typeInfo(ArgsType) != .Struct)
        @compileError("Expected tuple or struct, found " ++ @typeName(ArgsType));

    const fields_info = std.meta.fields(ArgsType);

    var answer: T = 0;

    comptime {
        var i: usize = 0;

        inline while (i < fields_info.len) : (i += 1) {
            switch (type_info) {
                .Int => answer +%= @field(args, fields_info[i].name),
                .Float => answer += @field(args, fields_info[i].name),
                else => unreachable,
            }
        }
    }

    return answer;
}