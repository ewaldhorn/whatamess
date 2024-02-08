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

test "Decode into a struct from a JSON string" {
    const options = std.json.ParseOptions{};

    // On success, `foo` will be of type `Data` struct.
    const foo = try std.json.parse(Data, &std.json.TokenStream.init(
        \\{
        \\    "name": "John Smith",
        \\    "is_alive": true,
        \\    "age": 45,
        \\    "address": {
        \\        "street_address": "21 2nd Street",
        \\        "city": "New York",
        \\        "state": "NY",
        \\        "postal_code": "10021-3100"
        \\    },
        \\    "children_age": [12, 8, 21],
        \\    "spouse": null
        \\}
    ), options);

    defer std.json.parseFree(Data, foo, options);

    try std.testing.expectEqualStrings(foo.name, "John Smith");
    try std.testing.expectEqual(foo.is_alive, true);
    try std.testing.expectEqual(foo.age, 45);
    try std.testing.expectEqualStrings(foo.address.street_address, "21 2nd Street");
    try std.testing.expectEqualSlices(i8, foo.children_age, &.{ 12, 8, 21 });
    try std.testing.expect(foo.spouse == null);
}
