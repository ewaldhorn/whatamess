const std = @import("std");

pub fn hello() []const u8 {
    const str = "Hello, World!";
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
	const allocator = gpa.allocator();    

    var list = std.ArrayList(u8).initCapacity(allocator, str.len) catch |err| {
        if (err == std.mem.Allocator.Error.OutOfMemory) {
            return "";
        } else {
            return "";
        }
    };
    
    defer list.deinit();

    for (str) |c| {
        list.appendAssumeCapacity(c);
    }

    const result = list.toOwnedSlice() catch |err| {
        if (err == std.mem.Allocator.Error.OutOfMemory) {
            return "";
        } else {
            return "";
        }
    };
    return result;
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 

test "can return a string" {
    const tmp = hello();

    try std.testing.expect(std.mem.eql(u8, tmp, "Hello, World!"));
}