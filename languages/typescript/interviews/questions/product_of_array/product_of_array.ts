// ------------------------------------------------------------------------- productExceptSelfBasic
export function productExceptSelfBasic(numbers: number[]): number[] {
  const count = numbers.length;
  const result: number[] = new Array();

  for (let outer = 0; outer < count; outer++) {
    let tmp = 1;

    for (let inner = 0; inner < count; inner++) {
      if (inner != outer) {
        tmp *= numbers[inner];
      }
    }

    // because JavaScript!
    if (tmp == -0) {
      tmp = 0;
    }

    result.push(tmp);
  }

  return result;
}

// ---------------------------------------------------------------------- productExceptSelfImproved
export function productExceptSelfImproved(numbers: number[]): number[] {
  const count = numbers.length;
  const result = new Array<number>(count);

  // work from left to right
  for (let i = 0, left = 1; i < count; ++i) {
    result[i] = left;
    left *= numbers[i];
  }

  // now from right to left
  for (let i = count - 1, right = 1; i >= 0; --i) {
    result[i] *= right;
    right *= numbers[i];
  }

  // because JavaScript!
  for (let i = 0; i < count; i++) {
    if (result[i] == -0) {
      result[i] = 0;
    }
  }

  return result;
}

// ------------------------------------------------------------------------ productExceptSelfMapped
export function productExceptSelfMapped(numbers: number[]): number[] {
  return numbers
    .map((_, i) =>
      numbers.reduce((pre, val, j) => pre * (i === j ? 1 : val), 1),
    )
    .map((v) => {
      if (v == -0) {
        return 0;
      } else {
        return v;
      }
    });
}
