pub const Logger = struct {
	level: Level,

	// "error" is reserved, names inside an @"..." are always
	// treated as identifiers
	const Level = enum {
		debug,
		info,
		@"error",
		fatal,
	};

	fn info(logger: Logger, msg: []const u8, out: anytype) !void {
		if (@intFromEnum(logger.level) <= @intFromEnum(Level.info)) {
			try out.writeAll(msg);
		}
	}
};