def find_string(nums: list[str], length: int) -> str:
    for n in all_possible_strings(length):
        if n not in nums:
            return n
    return "--"


# generate all the possible binary values for this length
def all_possible_strings(length: int) -> list[str]:
    possible_strings = []
    max_value = pow(2, length)
    for i in range(max_value):
        binary = bin(i)[2:]
        binary = "0" * (length - len(binary)) + binary
        possible_strings.append(binary)
    return possible_strings


# Check if the input value is in the answers list
def list_has(answers: list[str], input: str) -> bool:
    if input in answers:
        return True
    return False


# ======================================================================= TESTS
two_result_one = find_string(["00", "01"], 2)
assert list_has(["11", "10"], two_result_one)

two_result_two = find_string(["10", "01"], 2)
assert list_has(["00", "11"], two_result_two)

three_result_one = find_string(["111", "011", "001"], 3)
assert list_has(["000", "010", "100", "101", "110"], three_result_one)
