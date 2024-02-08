const std = @import("std");

pub fn isValid(s: []const u8) bool {
    if (s.len > 1) {
        // only bother if we have a potential number
        var total: u64 = 0;

        var pos: u8 = 1; // need to start from the right, working backwards, ignoring spaces
        var idx = s.len; // need to work through all characters, including spaces
        var digits: u8 = 0;

        while (idx != 0) {
            idx -= 1;
            const c = s[idx];
            if ((c < '0' or c > '9') and c != ' ') {
                return false;
            }

            if (c != ' ') {
                if (pos % 2 == 0) {
                    total += if ((c - '0') * 2 > 9) ((c - '0') * 2) - 9 else (c - '0') * 2;
                } else {
                    total += c - '0';
                }
                // we have a digit
                digits += 1;

                // increment our digit position
                pos += 1;
            }
        }

        return total % 10 == 0 and digits > 1;
    } else {
        return false;
    }
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const testing = std.testing;

test "single digit strings cannot be valid" {
    try testing.expect(!isValid("1"));
}

test "a single zero is invalid" {
    try testing.expect(!isValid("0"));
}

test "a simple valid SIN that remains valid if reversed" {
    try testing.expect(isValid("059"));
}

test "a simple valid SIN that becomes invalid if reversed" {
    try testing.expect(isValid("59"));
}

test "a valid Canadian SIN" {
    try testing.expect(isValid("055 444 285"));
}

test "invalid Canadian SIN" {
    try testing.expect(!isValid("055 444 286"));
}

test "invalid credit card" {
    try testing.expect(!isValid("8273 1232 7352 0569"));
}

test "invalid long number with an even remainder" {
    try testing.expect(!isValid("1 2345 6789 1234 5678 9012"));
}

test "invalid long number with a remainder divisible by 5" {
    try testing.expect(!isValid("1 2345 6789 1234 5678 9013"));
}

test "valid number with an even number of digits" {
    try testing.expect(isValid("095 245 88"));
}

test "valid number with an odd number of spaces" {
    try testing.expect(isValid("234 567 891 234"));
}

test "valid strings with a non-digit added at the end become invalid" {
    try testing.expect(!isValid("059a"));
}

test "valid strings with punctuation included become invalid" {
    try testing.expect(!isValid("055-444-285"));
}

test "valid strings with symbols included become invalid" {
    try testing.expect(!isValid("055# 444$ 285"));
}

test "single zero with space is invalid" {
    try testing.expect(!isValid(" 0"));
}

test "more than a single zero is valid" {
    try testing.expect(isValid("0000 0"));
}

test "input digit 9 is correctly converted to output digit 9" {
    try testing.expect(isValid("091"));
}

test "very long input is valid" {
    try testing.expect(isValid("9999999999 9999999999 9999999999 9999999999"));
}

test "valid luhn with an odd number of digits and non zero first digit" {
    try testing.expect(isValid("109"));
}

test "using ascii value for non-doubled non-digit isn't allowed" {
    try testing.expect(!isValid("055b 444 285"));
}

test "using ascii value for doubled non-digit isn't allowed" {
    try testing.expect(!isValid(":9"));
}

test "non-numeric, non-space char in the middle with a sum that's divisible by 10 isn't allowed" {
    try testing.expect(!isValid("59%59"));
}
