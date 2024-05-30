import { expect, test, describe } from "bun:test";
import { productExceptSelfBasic, productExceptSelfImproved, productExceptSelfMapped } from "./product_of_array";

describe("product of array", () => {
  test("[1,2,3,4]", () => {
    let expected = [24, 12, 8, 6];
    let result = productExceptSelfBasic([1, 2, 3, 4]);
    expect(result).toEqual(expected);
  });

  test("[-1, 1, 0, -3, 3]", () => {
    let expected = [0, 0, 9, 0, 0];
    expect(productExceptSelfBasic([-1, 1, 0, -3, 3])).toEqual(expected);
  });
});

describe("product of array improved", () => {
  test("[1,2,3,4]", () => {
    let expected = [24, 12, 8, 6];
    let result = productExceptSelfImproved([1, 2, 3, 4]);
    expect(result).toEqual(expected);
  });

  test("[-1, 1, 0, -3, 3]", () => {
    let expected = [0, 0, 9, 0, 0];
    expect(productExceptSelfImproved([-1, 1, 0, -3, 3])).toEqual(expected);
  });
});

describe("product of array using map", () => {
  test("[1,2,3,4]", () => {
    let expected = [24, 12, 8, 6];
    let result = productExceptSelfMapped([1, 2, 3, 4]);
    expect(result).toEqual(expected);
  });

  test("[-1, 1, 0, -3, 3]", () => {
    let expected = [0, 0, 9, 0, 0];
    expect(productExceptSelfMapped([-1, 1, 0, -3, 3])).toEqual(expected);
  });
});
