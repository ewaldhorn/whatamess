const std = @import("std");

const webui = @import("webui");

var rand: ?std.rand.Random = null;

pub fn main() !void {
    var nwin = webui.newWindow();

    _ = nwin.setRootFolder("web");

    try initRandomNumberGenerator();

    // bind to the button
    _ = nwin.bind("getTargetButton", get_target_button);

    // we can also use bind to make Zig functions callable by the JS side of things
    _ = nwin.bind("GetRandomNumber", returnRandomNumber);

    // show embedded index file
    _ = nwin.show("index.html");

    webui.wait();
    webui.clean();
}

fn initRandomNumberGenerator() !void {
    var prng = std.rand.DefaultPrng.init(blk: {
        var seed: u64 = undefined;
        try std.os.getrandom(std.mem.asBytes(&seed));
        break :blk seed;
    });
    rand = prng.random();
}

// returns a random number to webui
fn returnRandomNumber(e: webui.Event) void {
    webui.returnInt(e, getRandomNumber());
}

fn getRandomNumber() u8 {
    if (rand) |r| {
        return r.intRangeAtMost(u8, 1, 100);
    } else {
        return 0;
    }
}

fn get_target_button(e: webui.Event) void {
    var new_e = e;
    var win = new_e.getWindow();

    var js: [64]u8 = std.mem.zeroes([64]u8);
    const buf = std.fmt.bufPrint(&js, "setTarget({});", .{getRandomNumber()}) catch unreachable;

    win.run(buf);
}
