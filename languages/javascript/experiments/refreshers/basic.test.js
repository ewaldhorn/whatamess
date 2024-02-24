import { expect, test } from "bun:test";
import { returnTwo } from "./basic.js";

test("test returnTwo", () => {
    expect(returnTwo()).toBe(2);
    expect(returnTwo()).not.toBe(4);
});
