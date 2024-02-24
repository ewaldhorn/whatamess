const std = @import("std");

pub fn foo(iface: *Interface) void {
    var i: usize = 1;
    while (i < 4) : (i += 1) {
        const p = iface.pick();
        std.debug.print("foo {d}: {d}\n", .{ i, p });
    }
}

const Interface = struct {
    // can call directly: iface.pickFn(iface)
    pickFn: fn (*Interface) i32,

    // allows calling: iface.pick()
    pub fn pick(iface: *Interface) i32 {
        return iface.pickFn(iface);
    }
};

const PickRandom = struct {
    // specific to PickRandom
    r: std.rand.DefaultPrng,

    // implement the interface
    interface: Interface,

    fn init() PickRandom {
        return .{
            .r = std.rand.DefaultPrng.init(0),
            // point the interface function pointer to our function
            .interface = Interface{ .pickFn = myPick },
        };
    }

    fn myPick(iface: *Interface) i32 {
        // compute pointer to PickRandom struct from interface member pointer
        const self = @fieldParentPtr(PickRandom, "interface", iface);
        return self.r.random.intRangeAtMost(i32, 10, 20);
    }
};

const PickSeq = struct {
    x: i32,

    interface: Interface,

    fn init() PickSeq {
        return .{
            .x = 100,
            .interface = Interface{ .pickFn = myPick },
        };
    }

    fn myPick(iface: *Interface) i32 {
        const self = @fieldParentPtr(PickSeq, "interface", iface);
        self.x += 1;
        return self.x;
    }
};

pub fn main() !void {
    var pick_random = PickRandom.init();
    foo(&pick_random.interface);
    std.debug.print("main: {d}\n", .{pick_random.interface.pick()});

    var pick_seq = PickSeq.init();
    foo(&pick_seq.interface);
    std.debug.print("main: {d}\n", .{pick_seq.interface.pick()});

    // example of how copying an interface is wrong
    // var junk: [100]u8 = [_]u8{200} ** 100;
    var iface_copy = pick_seq.interface;
    // this will output junk numbers
    foo(&iface_copy);
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
