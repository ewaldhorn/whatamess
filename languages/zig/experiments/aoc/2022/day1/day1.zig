const std = @import("std");
const inputFile = @embedFile("./input.txt");

pub fn main() !void {
    std.debug.print("AOC 2022 Day 1 Part 1\n", .{});

    std.debug.print("The highest value is {d}.\n", .{try getHighest(inputFile)});
}

// Take input data (array of characters, split into lines and find the highest sum of continous numbers)
pub fn getHighest(data: []const u8) !u32 {
    var lines = std.mem.split(u8, data, "\n");
    var current : u32 = 0;
    var highest : u32 = 0;

    while (lines.next()) |num| {
        if (num.len == 0) {
            if (current > highest) {
                highest = current;
            }
            current = 0;
        } else {
        const tmp = try std.fmt.parseInt(u32, num,10);
        current += tmp;
        }
    }

    return highest;
}