const std = @import("std");

const Monster = struct {
    element: enum { fire, water, earth, wind },
    hp: u32,
};

const MonsterList = std.MultiArrayList(Monster);

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    var soa = MonsterList{};
    defer soa.deinit(allocator);

    // Add a couple of monsters
    for (0..10) |_| {
        try soa.append(allocator, .{
            .element = .fire,
            .hp = 20,
        });
    }

    // Count the number of fire monsters
    var total_fire: usize = 0;
    for (soa.items(.element)) |t| {
        if (t == .fire) total_fire += 1;
    }

    // Heal all monsters
    for (soa.items(.hp)) |*hp| {
        hp.* = 100;
    }

    std.debug.print("We have {d} fire monsters.\n", .{total_fire});
}
