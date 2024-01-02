const std = @import("std");
const expect = std.testing.expect;

const Number = union {
	int: i64,
	float: f64,
	nan: void,
};

test "setting nan" {
    const a = Number{.nan={}};

    try expect(a.nan == {});
}

test "setting float" {
    const a = Number{.float = 123.45};

    try expect(a.float == 123.45);
}