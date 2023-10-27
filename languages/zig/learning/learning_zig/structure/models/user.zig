pub const User = struct {
    power: u64,
    name: []const u8,
};

// Struct with default values and containing another struct
pub const SuperUser = struct {
    user: User,
    isAdmin: bool = true,
};
