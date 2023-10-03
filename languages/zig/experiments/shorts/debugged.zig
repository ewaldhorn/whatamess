const std = @import("std");
const assert = std.debug.assert;

test "meta programming" {
    const data: Data = .{
        .count_donations = 12.34,
        .count_happy_people = 100,
        .count_balance = -1,
        .bike_flag = true
    };
    try std.testing.expectEqual(countStuff(data), 111.34000015258789);
}

const Data = struct {
    count_donations: f32,
    count_happy_people: u32,
    count_balance: i32,
    bike_flag: bool
};

fn countStuff(data: Data) f64 {
    var accumulator: f64 = 0;

    inline for (@typeInfo(Data).Struct.fields) |field| {
        if (std.mem.startsWith(u8, field.name, "count_")) {
            switch(field.type){
                f32 => accumulator += @field(data, field.name),
                u32 => accumulator += @floatFromInt(@field(data, field.name)),
                i32 => accumulator += @floatFromInt(@field(data, field.name)),
                else => {}
            }
        }
    }
    return accumulator;
}