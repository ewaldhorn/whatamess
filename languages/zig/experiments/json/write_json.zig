const std = @import("std");

const Data = struct {
    name: []const u8,
    is_alive: bool,
    age: u8,
    address: struct {
        street_address: []const u8,
        city: []const u8,
        state: []const u8,
        postal_code: []const u8,
    },
    children_age: []const i8,
    spouse: ?[]const u8,
};

test "Encode a struct into a JSON string" {
    const foo = Data{
        .name = "John Smith",
        .is_alive = true,
        .age = 45,
        .address = .{
            .street_address = "21 2nd Street",
            .city = "New York",
            .state = "NY",
            .postal_code = "10021-3100",
        },
        .children_age = &.{ 12, 8, 21 },
        .spouse = null,
    };

    var file = try std.fs.cwd().createFile("data.json", .{});
    defer file.close();

    const options = std.json.StringifyOptions{ .whitespace = .indent_2 };

    try std.json.stringify(foo, options, file.writer());

    _ = try file.write("\n");
}
