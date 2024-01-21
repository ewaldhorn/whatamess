const std = @import("std");
const expect = std.testing.expect;
const eql = std.mem.eql;

const ArrayList = std.ArrayList;
const test_allocator = std.testing.allocator;

// std.io.Writer and std.io.Reader provide standard ways of making use of IO. std.ArrayList(u8) has 
// a writer method which gives us a writer. Let's use it.
test "io writer usage" {
    var list = ArrayList(u8).init(test_allocator);
    defer list.deinit();

    const bytes_written = try list.writer().write(
        "Hello World!",
    );
    
    try expect(bytes_written == 12);
    try expect(eql(u8, list.items, "Hello World!"));
}