const std = @import("std");

pub const InventoryItem = struct {
    price: f32 = 0.00,
    name: []const u8,

    pub fn init(name: []const u8, price: f32) InventoryItem {
        return InventoryItem{
            .price = price,
            .name = name,
        };
    }

    pub fn report(item: InventoryItem) void {
        std.debug.print("Item '{s}' costs R{d:.2}.\n", .{ item.name, item.price });
    }
};
