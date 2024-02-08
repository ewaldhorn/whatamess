const std = @import("std");

// factory type
const Gpa = std.heap.GeneralPurposeAllocator(.{});

pub fn main() !void {
    // instantiates the factory
    var gpa = Gpa{};
    
    // retrieves the created allocator.
    var galloc = &gpa.allocator();
    
    // scopes the lifetime of the allocator to this function and
    // performs cleanup; 
    defer _ = gpa.deinit();

    var slice = try galloc.alloc(i32, 2);
    // comment to trigger memory leak warning
    defer galloc.free(slice);
    
    const single = try galloc.create(i32);
    // comment to trigger memory leak warning
    defer galloc.destroy(single);

    slice[0] = 47;
    slice[1] = 48;
    single.* = 49;

    std.debug.print("slice: [{}, {}]\n", .{slice[0], slice[1]});
    std.debug.print("single: {}\n", .{single.*});
}