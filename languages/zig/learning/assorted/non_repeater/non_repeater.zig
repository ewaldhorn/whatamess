const std = @import("std");

// Find the first occurence of a non-repeat sequence of four characters.
// This will be four unique characters, where none of the characters are the same.

pub fn main() void {
    const input = "znrznnozzroronznzon";
    std.debug.print("Examining:\n\"{s}\"\n\n", .{input});

    var pos : usize = 0;

    while (pos <= (input.len-4)) : (pos += 1) {
        std.debug.print("{d:3} {s}", .{pos, input[pos..pos+4]});    

        std.debug.print("\n", .{});
    }
}