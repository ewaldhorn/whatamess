const std = @import("std");

const source_file = @embedFile("./text.txt");

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
test "source is available" {
    var reader = std.mem.tokenizeAny(u8, source_file, "\n");
    
    std.debug.print("\n========================================================\n", .{});
    std.debug.print("Parsing embedded 'text.txt' file:\n\n", .{});
    while (reader.next()) |line| {
        std.debug.print("> {s}\n", .{line});
    }
    std.debug.print("\n", .{});
}