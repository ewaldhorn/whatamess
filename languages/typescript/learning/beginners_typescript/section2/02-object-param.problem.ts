import { expect, it } from "vitest";

type TwoNumberArgs = {
  first: number; second: number;
};

export const addTwoNumbersTyped = (params: TwoNumberArgs) => {
  return params.first + params.second;
}

export const addTwoNumbers = (params: { first: number; second: number; }) => {
  return params.first + params.second;
};

it("Should add the two numbers together", () => {
  expect(
    addTwoNumbers({
      first: 2,
      second: 4,
    }),
  ).toEqual(6);

  expect(
    addTwoNumbers({
      first: 10,
      second: 20,
    }),
  ).toEqual(30);
});
