pub fn squareOfSum(number: usize) usize {
    var result: usize = 0;

    var n: usize = 0;
    while (n <= number) : (n += 1) {
        result += n;
    }
    return result * result;
}

pub fn sumOfSquares(number: usize) usize {
    var result: usize = 0;

    for (1..(number + 1)) |n| {
        result += (n * n);
    }

    return result;
}

pub fn differenceOfSquares(number: usize) usize {
    var result = squareOfSum(number) - sumOfSquares(number);

    return if (result < 0) result * -1 else result;
}
