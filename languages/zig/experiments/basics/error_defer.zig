const std = @import("std");
const Allocator = std.mem.Allocator;

pub const Game = struct {
	players: []Player,
	history: []Move,
	allocator: Allocator,

	fn init(allocator: Allocator, player_count: usize) !Game {
		var players = try allocator.alloc(Player, player_count);
		errdefer allocator.free(players);

		// store 10 most recent moves per player
		var history = try allocator.alloc(Move, player_count * 10);

		return .{
			.players = players,
			.history = history,
			.allocator = allocator,
		};
	}

	fn deinit(game: Game) void {
		const allocator = game.allocator;
		allocator.free(game.players);
		allocator.free(game.history);
	}
};