def find_string(nums, length):
    return "00"


# Check if the input value is in the answers list
def list_has(answers, input):
    if input in answers:
        return True
    return False


# ======================================================================= TESTS
two_result = find_string(["00", "01"], 2)
assert list_has(["11", "10"], two_result)
