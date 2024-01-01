const std = @import("std");
const fs = std.fs;

pub fn main() !void {
    const file_to_print = "text.txt";

    // Defining this type made it easier for me to understand allocating a string array.
    const String = []const u8;

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();

    const allocator = gpa.allocator();

    // See also std.fs.selfExeDirPathAlloc.
    const exe_path = try std.fs.selfExePathAlloc(allocator);
    defer allocator.free(exe_path);
    const opt_exe_dir = std.fs.path.dirname(exe_path);

    if (opt_exe_dir) |exe_dir| {
        var paths = [_]String{ exe_dir, file_to_print };

        const file_path = try fs.path.join(allocator, &paths);
        defer allocator.free(file_path);
        std.debug.print("File location: {s}.\n", .{file_path});

        // There are a few ways to load the file (iterate over lines, etc).
        // Here we load the whole thing and print it all out:
        const file_content = std.fs.cwd().readFileAlloc(allocator, file_path, std.math.maxInt(usize)) catch {
            std.debug.print("Failed: {s} \n", .{file_to_print});
            return;
        };
        defer allocator.free(file_content);
        std.debug.print("{s}\n", .{file_content});
    }
}