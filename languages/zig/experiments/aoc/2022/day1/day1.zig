const std = @import("std");
const inputFile = @embedFile("./input.txt");

pub fn main() !void {
    std.debug.print("AOC 2022 Day 1 Part 1\n", .{});

    std.debug.print("The highest value is {d}. (Expect 24000)\n", .{try getHighest(inputFile)});
    std.debug.print("The sum of the highest three values is {d}. (Expect 45000)\n", .{try getHighestThree(inputFile)});
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

pub fn getHighestThree(data: []const u8) !u32 {
    var lines = std.mem.split(u8, data, "\n");
    var arr = [_]u32{0} ** 3;
    var current : u32 = 0;

    while (lines.next()) |num| {
        if (num.len == 0) {
            const i = min(&arr);
            if (current > arr[i]) {
                arr[i] = current;
            }
            current = 0;
        } else {
        const tmp = try std.fmt.parseInt(u32, num,10);
        current += tmp;
        }
    }

    return sum(&arr);
}

// Finds the smallest number in the array and returns its position
pub fn min(arr: []u32) usize {
    var result: usize = 0;

    for (0..3) |i| {
        if (arr[i] < arr[result]) {
            result = i;
        }
    }

    return result;
}

// Tallies up our array of top results into one number
pub fn sum(arr: []u32) u32 {
    var result: u32 = 0;

    for (arr) |element| {
        result += element;
    }

    return result;
}