fun main() {
    var input = arrayOf<Int>(1, 2, 3, 4)
    var expected = arrayOf(24, 12, 8, 6)

    var result = productExceptSelfBasic(input)
    println("Content matched: ${expected.contentEquals(result)}")

    input = arrayOf(-1, 1, 0, -3, 3)
    expected = arrayOf(0, 0, 9, 0, 0)
    result = productExceptSelfBasic(input)
    println("Content matched: ${expected.contentEquals(result)}")

    result = productExceptSelfBetter(input)
    println("Content matched: ${expected.contentEquals(result)}")
}

// ------------------------------------------------------------------------- productExceptSelfBasic
fun productExceptSelfBasic(numbers: Array<Int>): Array<Int> {
    val result = IntArray(numbers.size)

    for (outer in numbers.indices) {
        var tmp = 1

        for (inner in numbers.indices) {
            if (outer != inner) {
                tmp = tmp * numbers[inner]
            }
        }

        result[outer] = tmp
    }

    return result.toTypedArray()
}

// ------------------------------------------------------------------------ productExceptSelfBetter
fun productExceptSelfBetter(numbers: Array<Int>): Array<Int> {
    val result = IntArray(numbers.size) { 1 }

    var tmp = 1
    for (left in 0..numbers.size - 1) {
        result[left] = tmp
        tmp *= numbers[left]
    }

    tmp = 1
    for (right in (numbers.size - 1) downTo 0) {
        result[right] *= tmp
        tmp *= numbers[right]
    }

    return result.toTypedArray()
}
