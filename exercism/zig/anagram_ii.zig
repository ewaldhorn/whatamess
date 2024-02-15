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

    for (candidates) |candidate| {
        if (candidate.len != word.len) { // skip impossible combinations
            continue;
        }

        if (isSameWord(candidate, word)) { // skip same words
            continue;
        }

        if (hasSameFrequencies(candidate, word)) {
            try results.insert(candidate);
        }
    }

    return results;
}

// Checks if two words have the same letter frequencies
fn hasSameFrequencies(left: []const u8, right: []const u8) bool {
    if (left.len != right.len) { // guard against impossible combinations
        return false;
    }

    for (left) |c| {
        if (countLetterOccurences(left, c) != countLetterOccurences(right, c)) {
            return false;
        }
    }

    return true;
}

// Count how many times a letter occurs in a word
fn countLetterOccurences(source: []const u8, target: u8) u8 {
    var count: u8 = 0;

    for (source) |letter| {
        if (std.ascii.toLower(letter) == std.ascii.toLower(target)) {
            count += 1;
        }
    }

    return count;
}

// Check if two words are the same, ignoring case
fn isSameWord(left: []const u8, right: []const u8) bool {
    if (left.len != right.len) {
        return false;
    }

    for (left, right) |l, r| {
        if (std.ascii.toLower(l) != std.ascii.toLower(r)) {
            return false;
        }
    }

    return true;
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
