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
