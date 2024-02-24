import { describe, expect, test } from "bun:test";
import { noDuplicates, noDuplicatesAlternative, noDuplicatesSorted } from "./no_duplicates.js";

// test basic solution

describe("basic", () => {
    test("test empty string", () => {
        expect(noDuplicates("")).toBe(true);
    });

    test("test no duplicates", () => {
        expect(noDuplicates("abcd")).toBe(true);
    });

    test("test has duplicate", () => {
        expect(noDuplicates("abca")).toBe(false);
    });

    test("test has multiple duplicates", () => {
        expect(noDuplicates("abcdefghijklmnaopqrstub")).toBe(false);
    });

    test("test has long no duplicates", () => {
        expect(noDuplicates("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")).toBe(true);
    });

    test("test has long with duplicates", () => {
        expect(noDuplicates("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa")).toBe(false);
    });
});

// test alternative solution
describe("alternative", () => {
    test("test empty string - alternative", () => {
        expect(noDuplicatesAlternative("")).toBe(true);
    });

    test("test no duplicates - alternative", () => {
        expect(noDuplicatesAlternative("abcd")).toBe(true);
    });

    test("test has duplicate - alternative", () => {
        expect(noDuplicatesAlternative("abca")).toBe(false);
    });

    test("test has multiple duplicates - alternative", () => {
        expect(noDuplicatesAlternative("abcdefghijklmnaopqrstub")).toBe(false);
    });

    test("test has long no duplicates - alternative", () => {
        expect(noDuplicatesAlternative("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")).toBe(true);
    });

    test("test has long with duplicates - alternative", () => {
        expect(noDuplicatesAlternative("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa")).toBe(false);
    });
});

// test sorted solution
describe("sorted", () => {
    test("test empty string - sorted", () => {
        expect(noDuplicatesSorted("")).toBe(true);
    });

    test("test no duplicates - sorted", () => {
        expect(noDuplicatesSorted("abcd")).toBe(true);
    });

    test("test has duplicate - sorted", () => {
        expect(noDuplicatesSorted("abca")).toBe(false);
    });

    test("test has multiple duplicates - sorted", () => {
        expect(noDuplicatesSorted("abcdefghijklmnaopqrstub")).toBe(false);
    });

    test("test has long no duplicates - sorted", () => {
        expect(noDuplicatesSorted("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")).toBe(true);
    });

    test("test has long with duplicates - sorted", () => {
        expect(noDuplicatesSorted("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa")).toBe(false);
    });
});