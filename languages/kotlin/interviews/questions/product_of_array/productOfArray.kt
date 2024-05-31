fun main() {
    var input = arrayOf<Int>(1, 2, 3, 4)
    var expected = arrayOf(24, 12, 8, 6)

    var result = productExceptSelfBasic(input)
    println("Content matched: ${expected.contentEquals(result)}")

    input = arrayOf(-1, 1, 0, -3, 3)
    expected = arrayOf(0, 0, 9, 0, 0)
    result = productExceptSelfBasic(input)
    println("Content matched: ${expected.contentEquals(result)}")
}

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
