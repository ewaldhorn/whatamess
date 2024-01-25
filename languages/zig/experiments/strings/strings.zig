const assert = @import("std").debug.assert;
const mem = @import("std").mem;

test "string literals" {
    // In Zig a string literal is an array of bytes.
    const normal_bytes = "hello";
    assert(@TypeOf(normal_bytes) == [5]u8);
    assert(normal_bytes.len == 5);
    assert(normal_bytes[1] == 'e');
    assert('e' == '\x65');
    assert(mem.eql(u8, "hello", "h\x65llo"));

    // A C string literal is a null terminated pointer.
    const null_terminated_bytes = "hello";
    assert(@TypeOf(null_terminated_bytes) == u8);
    assert(null_terminated_bytes[5] == 0);
}