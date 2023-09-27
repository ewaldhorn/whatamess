// @ts-check

/**
 * Calculates the sum of the two input arrays.
 *
 * @param {number[]} array1
 * @param {number[]} array2
 * @returns {number} sum of the two arrays
 */
export function twoSum(array1, array2) {
  let num1 = Number(array1.join(""));
  let num2 = Number(array2.join(""));
  return num1 + num2;
}

/**
 * Checks whether a number is a palindrome.
 *
 * @param {number} value
 * @returns {boolean} whether the number is a palindrome or not
 */
export function luckyNumber(value) {
  let s = String(value);
  let s_reversed = s.split("").reverse().join("");

  return s == s_reversed;
}

/**
 * Determines the error message that should be shown to the user
 * for the given input value.
 *
 * @param {string|null|undefined} input
 * @returns {string} error message
 */
export function errorMessage(input) {
  if (input === undefined || input === null || input == "") {
    return "Required field";
  } else if (input === 0 || !Number(input)) {
    return "Must be a number besides 0";
  } else {
    return "";
  }
}

