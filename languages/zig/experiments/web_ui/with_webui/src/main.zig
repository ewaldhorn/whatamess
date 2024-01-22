const std = @import("std");

const webui = @import("webui");
const indexFile = @embedFile("index.html");

var rand: ?std.rand.Random = null;

pub fn main() !void {
    var nwin = webui.newWindow();

    var prng = std.rand.DefaultPrng.init(blk: {
        var seed: u64 = undefined;
        try std.os.getrandom(std.mem.asBytes(&seed));
        break :blk seed;
    });
    rand = prng.random();

    // show embedded index file
    _ = nwin.show(indexFile);

    // bind to the button
    _ = nwin.bind("getTargetButton", get_target_button);

    webui.wait();
    webui.clean();
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
