const std = @import("std");
const expect = std.testing.expect;

const Direction = enum { north, south, east, west };
const Value = enum(u2) { first, second, third };

test "enum ordinal value" {
    try expect(@intFromEnum(Value.first) == 0);
    try expect(@intFromEnum(Value.second) == 1);
    try expect(@intFromEnum(Value.third) == 2);
}

// enum values can be overriden, and will then continue from there
const Value2 = enum(u32) {
    hundred = 100,
    thousand = 1000,
    million = 1000000,
    next,
};

test "set enum ordinal value" {
    try expect(@intFromEnum(Value2.hundred) == 100);
    try expect(@intFromEnum(Value2.thousand) == 1000);
    try expect(@intFromEnum(Value2.million) == 1000000);
    try expect(@intFromEnum(Value2.next) == 1000001);
}

// enums can have functions associated with them
// these are then namespaced to the enum
const Suit = enum {
    clubs,
    spades,
    diamonds,
    hearts,
    pub fn isClubs(self: Suit) bool {
        return self == Suit.clubs;
    }
};

test "enum method" {
    try expect(Suit.spades.isClubs() == Suit.isClubs(.spades));
}

// enums can have var and const declarations, these are then globals to the enum and
// are not attached to particular instance of the enum

const Mode = enum {
    var count: u32 = 0;
    on,
    off,
};

test "hmm" {
    Mode.count += 1;
    try expect(Mode.count == 1);
}