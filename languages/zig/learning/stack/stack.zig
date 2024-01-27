const std = @import("std");
const ArrayList = std.ArrayList;
const Allocator = std.mem.Allocator;

// Generic stack
// Based on https://www.huy.rocks/everyday/01-08-2022-zig-case-study-implementing-a-generic-stack
// Updated to work with Zig 0.12.x

// Stack functions
// init(): Initialize the stack
// deinit(): Deinitialize the stack
// push(): Push a new item to the stack
// pop(): Remove and return the top item of the stack
// top(): Get the top item of the stack
// count(): Get the length of the items in the stack
// isEmpty(): Check if the stack is empty or not

pub fn Stack(comptime T: type) type {
    return struct {
        stack: ArrayList(T),

        const Self = @This();

        pub fn init(allocator: Allocator) Self {
            return Self{ .stack = ArrayList(T).init(allocator) };
        }

        pub fn deinit(self: *Self) void {
            self.stack.deinit();
        }

        pub fn push(self: *Self, val: T) !void {
            try self.stack.append(val);
        }

        pub fn pop(self: *Self) ?T {
            return self.stack.popOrNull();
        }

        pub fn top(self: *Self) ?T {
            if (self.stack.items.len == 0) {
                return null;
            }
            return self.stack.items[self.stack.items.len - 1];
        }

        pub fn count(self: *Self) usize {
            return self.stack.items.len;
        }

        pub fn isEmpty(self: *Self) bool {
            return self.count() == 0;
        }
    };
}

// ========================================================================================== TESTS

test {
    const expect = std.testing.expect;

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer std.debug.assert(gpa.deinit() == std.heap.Check.ok);

    const IntStack = Stack(i32);

    var stack = IntStack.init(gpa.allocator());
    defer stack.deinit();

    try stack.push(1);
    try stack.push(2);
    try stack.push(3);

    try expect(stack.isEmpty() == false);

    try expect(stack.top().? == 3);
    try expect(stack.pop().? == 3);
    try expect(stack.top().? == 2);
    try expect(stack.pop().? == 2);
    try expect(stack.top().? == 1);
    try expect(stack.pop().? == 1);

    try expect(stack.isEmpty() == true);
}
