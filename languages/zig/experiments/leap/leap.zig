const expect = @import("std").testing.expect;

pub fn isLeapYear(year: u32) bool {
    return (year % 4 == 0) and ((year % 100 != 0) or (year % 400 == 0));
}

test "year 1700" {
    const result = isLeapYear(1700);
    try expect(result == false);
}

test "year 1701" {
    const result = isLeapYear(1701);
    try expect(result == false);
}

test "year 1702" {
    const result = isLeapYear(1702);
    try expect(result == false);
}

test "year 1703" {
    const result = isLeapYear(1703);
    try expect(result == false);
}

test "year 1704" {
    const result = isLeapYear(1704);
    try expect(result == true);
}

test "year 1705" {
    const result = isLeapYear(1705);
    try expect(result == false);
}

test "year 1706" {
    const result = isLeapYear(1706);
    try expect(result == false);
}

test "year 1707" {
    const result = isLeapYear(1707);
    try expect(result == false);
}

test "year 1708" {
    const result = isLeapYear(1708);
    try expect(result == true);
}

test "year 1800" {
    const result = isLeapYear(1800);
    try expect(result == false);
}

test "year 1900" {
    const result = isLeapYear(1900);
    try expect(result == false);
}

test "year 2000" {
    const result = isLeapYear(2000);
    try expect(result == true);
}

test "year 2001" {
    const result = isLeapYear(2001);
    try expect(result == false);
}

test "year 2002" {
    try expect(isLeapYear(2002) == false);
}

test "year 2004" {
    const result = isLeapYear(2004);
    try expect(result == true);
}

test "year 2008" {
    const result = isLeapYear(2008);
    try expect(result == true);
}

test "year 2012" {
    const result = isLeapYear(2012);
    try expect(result == true);
}

test "year 2048" {
    const result = isLeapYear(2048);
    try expect(result == true);
}
