const std = @import("std");

pub fn isPangram(str: []const u8) bool {
    var counter = [_]u8{0} ** 26;
    
    
    if (str.len < 26) {
        // don't bother, this will never work
        return false;
    }

    const diff : u8 = 'a' - 'A';

    for (str) |c| {
        switch (c) {
            'a'...'z' => {
                counter[c - 'a'] += 1;
            },
            'A'...'Z' => {
                counter[c + diff - 'a'] += 1;
            },
            else => continue
        }
    }

    var answer = true;

    for (counter) |v| {
        answer = answer and (v >= 1);
    }

    return answer;
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
// A 65
// a 97 32
const testing = std.testing;

test "empty sentence" {
    try testing.expect(!isPangram(""));
}

test "perfect lower case" {
    try testing.expect(isPangram("abcdefghijklmnopqrstuvwxyz"));
}

test "only lower case" {
    try testing.expect(isPangram("the quick brown fox jumps over the lazy dog"));
}

test "missing the letter 'x'" {
    try testing.expect(!isPangram("a quick movement of the enemy will jeopardize five gunboats"));
}

test "missing the letter 'h'" {
    try testing.expect(!isPangram("five boxing wizards jump quickly at it"));
}

test "with underscores" {
    try testing.expect(isPangram("the_quick_brown_fox_jumps_over_the_lazy_dog"));
}

test "with numbers" {
    try testing.expect(isPangram("the 1 quick brown fox jumps over the 2 lazy dogs"));
}

test "missing letters replaced by numbers" {
    try testing.expect(!isPangram("7h3 qu1ck brown fox jumps ov3r 7h3 lazy dog"));
}

test "mixed case and punctuation" {
    try testing.expect(isPangram("\"Five quacking Zephyrs jolt my wax bed.\""));
}

test "a-m and A-M are 26 different characters but not a pangram" {
    try testing.expect(!isPangram("abcdefghijklm ABCDEFGHIJKLM"));
}

test "non-alphanumeric printable ASCII" {
    try testing.expect(!isPangram(" !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"));
}