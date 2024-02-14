const std = @import("std");
const mem = std.mem;

/// Returns the set of strings in `candidates` that are anagrams of `word`.
/// Caller owns the returned memory.
pub fn detectAnagrams(
    allocator: mem.Allocator,
    word: []const u8,
    candidates: []const []const u8,
) !std.BufSet {
    var results = std.BufSet.init(allocator); // allocate our bufset
    
    if (candidates.len == 0) {
        return results; // nothing to do, break now
    }

    const word_lowercase = try std.ascii.allocLowerString(allocator, word);
    defer allocator.free(word_lowercase);

    // first try to allocate a working buffer for the word so we can sort it
    const word_buffer = try allocator.alloc(u8, word.len);
    defer allocator.free(word_buffer);

    // now copy the word and sort the letters
    std.mem.copyForwards(u8, word_buffer, word_lowercase);
    std.mem.sort(u8, word_buffer, {}, comptime std.sort.asc(u8));

    for (0..candidates.len) |idx| {
        const candidate_lowercase = try std.ascii.allocLowerString(allocator, candidates[idx]);
        defer allocator.free(candidate_lowercase);

        if (!std.mem.eql(u8, word_lowercase, candidate_lowercase)) {
            // we know it's not the same word, so let's sort the letters
            const candidate_buffer = try allocator.alloc(u8, candidate_lowercase.len);
            defer allocator.free(candidate_buffer);

            // copy and sort the candidate letters
            std.mem.copyForwards(u8, candidate_buffer, candidate_lowercase);
            std.mem.sort(u8, candidate_buffer, {}, comptime std.sort.asc(u8));

            // if it's a match, add it
            if (std.mem.eql(u8, word_buffer, candidate_buffer)) {
                try results.insert(candidates[idx]);
            }
        }
    }

    return results;
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const testing = std.testing;

fn testAnagrams(
    expected: []const []const u8,
    word: []const u8,
    candidates: []const []const u8,
) !void {
    var actual = try detectAnagrams(testing.allocator, word, candidates);
    defer actual.deinit();
    try testing.expectEqual(expected.len, actual.count());
    for (expected) |e| {
        try testing.expect(actual.contains(e));
    }
}

test "no matches" {
    const expected = [_][]const u8{};
    const word = "diaper";
    const candidates = [_][]const u8{ "hello", "world", "zombies", "pants" };
    try testAnagrams(&expected, word, &candidates);
}

test "detects two anagrams" {
    const expected = [_][]const u8{ "lemons", "melons" };
    const word = "solemn";
    const candidates = [_][]const u8{ "lemons", "cherry", "melons" };
    try testAnagrams(&expected, word, &candidates);
}

test "does not detect anagram subsets" {
    const expected = [_][]const u8{};
    const word = "good";
    const candidates = [_][]const u8{ "dog", "goody" };
    try testAnagrams(&expected, word, &candidates);
}

test "detects anagram" {
    const expected = [_][]const u8{"inlets"};
    const word = "listen";
    const candidates = [_][]const u8{ "enlists", "google", "inlets", "banana" };
    try testAnagrams(&expected, word, &candidates);
}

test "detects three anagrams" {
    const expected = [_][]const u8{ "gallery", "regally", "largely" };
    const word = "allergy";
    const candidates = [_][]const u8{ "gallery", "ballerina", "regally", "clergy", "largely", "leading" };
    try testAnagrams(&expected, word, &candidates);
}

test "detects multiple anagrams with different case" {
    const expected = [_][]const u8{ "Eons", "ONES" };
    const word = "nose";
    const candidates = [_][]const u8{ "Eons", "ONES" };
    try testAnagrams(&expected, word, &candidates);
}

test "does not detect non-anagrams with identical checksum" {
    const expected = [_][]const u8{};
    const word = "mass";
    const candidates = [_][]const u8{"last"};
    try testAnagrams(&expected, word, &candidates);
}

test "detects anagrams case-insensitively" {
    const expected = [_][]const u8{"Carthorse"};
    const word = "Orchestra";
    const candidates = [_][]const u8{ "cashregister", "Carthorse", "radishes" };
    try testAnagrams(&expected, word, &candidates);
}
test "detects anagrams using case-insensitive subject" {
    const expected = [_][]const u8{"carthorse"};
    const word = "Orchestra";
    const candidates = [_][]const u8{ "cashregister", "carthorse", "radishes" };
    try testAnagrams(&expected, word, &candidates);
}

test "detects anagrams using case-insensitive possible matches" {
    const expected = [_][]const u8{"Carthorse"};
    const word = "orchestra";
    const candidates = [_][]const u8{ "cashregister", "Carthorse", "radishes" };
    try testAnagrams(&expected, word, &candidates);
}

test "does not detect an anagram if the original word is repeated" {
    const expected = [_][]const u8{};
    const word = "go";
    const candidates = [_][]const u8{"goGoGO"};
    try testAnagrams(&expected, word, &candidates);
}

test "anagrams must use all letters exactly once" {
    const expected = [_][]const u8{};
    const word = "tapper";
    const candidates = [_][]const u8{"patter"};
    try testAnagrams(&expected, word, &candidates);
}

test "words are not anagrams of themselves" {
    const expected = [_][]const u8{};
    const word = "BANANA";
    const candidates = [_][]const u8{"BANANA"};
    try testAnagrams(&expected, word, &candidates);
}
test "words are not anagrams of themselves even if letter case is partially different" {
    const expected = [_][]const u8{};
    const word = "BANANA";
    const candidates = [_][]const u8{"Banana"};
    try testAnagrams(&expected, word, &candidates);
}

test "words are not anagrams of themselves even if letter case is completely different" {
    const expected = [_][]const u8{};
    const word = "BANANA";
    const candidates = [_][]const u8{"banana"};
    try testAnagrams(&expected, word, &candidates);
}

test "words other than themselves can be anagrams" {
    const expected = [_][]const u8{"Silent"};
    const word = "LISTEN";
    const candidates = [_][]const u8{ "LISTEN", "Silent" };
    try testAnagrams(&expected, word, &candidates);
}
