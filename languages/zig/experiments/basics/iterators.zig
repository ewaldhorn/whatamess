const std = @import("std");
const expect = std.testing.expect;
const eql = std.mem.eql;

// It is a common idiom to have a struct type with a next function with an optional in its return type, 
// so that the function may return a null to indicate that iteration is finished.
// std.mem.SplitIterator (and the subtly different std.mem.TokenIterator) is an example of this pattern.
test "split iterator" {
    const text = "robust, optimal, reusable, maintainable, ";
    var iter = std.mem.split(u8, text, ", ");
    
    try expect(eql(u8, iter.next().?, "robust"));
    try expect(eql(u8, iter.next().?, "optimal"));
    try expect(eql(u8, iter.next().?, "reusable"));
    try expect(eql(u8, iter.next().?, "maintainable"));
    try expect(eql(u8, iter.next().?, ""));

    try expect(iter.next() == null);
}

// Some iterators have a !?T return type, as opposed to ?T. !?T requires that we unpack the error 
// union before the optional, meaning that the work done to get to the next iteration may error. 
// Here is an example of doing this with a loop. cwd has to be opened with iterate permissions in 
// order for the directory iterator to work.
test "iterator looping" {
    var iter = (try std.fs.cwd().openDir(
        ".",
        .{},
    )).iterate();

    var file_count: usize = 0;
    while (try iter.next()) |entry| {
        if (entry.kind == .file) file_count += 1;
    }

    try expect(file_count > 0);
}

// Here we will implement a custom iterator. This will iterate over a slice of strings, 
// yielding the strings which contain a given string.
const ContainsIterator = struct {
    strings: []const []const u8,
    needle: []const u8,
    index: usize = 0,
    fn next(self: *ContainsIterator) ?[]const u8 {
        const index = self.index;
        for (self.strings[index..]) |string| {
            self.index += 1;
            // return string if it has an index (was found), otherwise continue for loop
            if (std.mem.indexOf(u8, string, self.needle)) |_| {
                return string;
            }
        }
        return null;
    }
};

test "custom iterator" {
    var iter = ContainsIterator{
        .strings = &[_][]const u8{ "one", "two", "three" },
        .needle = "e",
    };

    try expect(eql(u8, iter.next().?, "one"));
    try expect(eql(u8, iter.next().?, "three"));
    try expect(iter.next() == null);
}