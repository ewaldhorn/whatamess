const std = @import("std");

fn validateCardNumber(card: [16]u8) bool {
    var grand_total : u32 = 0;

    var position = card.len;

    while (position > 0){
        position -= 1;
        if (position % 2 == 0) {
            const tmp = card[position] * 2;
            grand_total += if (tmp >= 10) tmp-9 else tmp;
        } else {
            grand_total += card[position];
        }
    }

    return grand_total % 10 == 0;
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const expect = std.testing.expect;

test "all zeroes" {
    const input = [16]u8{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0};

    try expect(validateCardNumber(input));
}

test "valid card" {
    const input = [16]u8{4,5,5,6,7,3,7,5,8,6,8,9,9,8,5,5};
    try expect(validateCardNumber(input));
}

test "invalid card" {
    const input = [16]u8{3,5,5,6,7,3,7,5,8,6,8,9,9,8,5,5};
    try expect(!validateCardNumber(input));
}