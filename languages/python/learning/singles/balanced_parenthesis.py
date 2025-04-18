# Problem Statement:
#   Given a string of parentheses, write a function that returns True if the parentheses are balanced and False otherwise.
# Example Inputs and Outputs:
#   balanced_parentheses("()") → True
#   balanced_parentheses("(()") → False
#   balanced_parentheses("())") → False
#   balanced_parentheses("((()))") → True
# Constraints:
#   The input string will only contain parentheses (i.e., ( and )).
#   The length of the input string will not exceed 100 characters.


def is_odd(n):
    return n % 2 != 0


def balanced_parentheses(input):
    if len(input) == 0:
        return True

    if is_odd(len(input)):
        return False

    balance = 0
    for char in input:
        if char == "(":
            balance += 1
        elif char == ")":
            if balance == 0:
                return False
            balance -= 1
    return balance == 0


# tests
import unittest


class TestBalancedParentheses(unittest.TestCase):
    def test_valid_balanced(self):
        self.assertTrue(balanced_parentheses("()"))
        self.assertTrue(balanced_parentheses("((()))"))
        self.assertTrue(balanced_parentheses("(()())"))
        self.assertTrue(balanced_parentheses("((((((()))))))"))

    def test_invalid_unbalanced(self):
        self.assertFalse(balanced_parentheses("(()"))
        self.assertFalse(balanced_parentheses("())"))
        self.assertFalse(balanced_parentheses("((())"))

    def test_edge_cases(self):
        self.assertTrue(balanced_parentheses(""))
        self.assertFalse(balanced_parentheses(")"))
        self.assertFalse(balanced_parentheses("("))
        self.assertFalse(balanced_parentheses("(("))
        self.assertFalse(balanced_parentheses("))"))
        self.assertFalse(balanced_parentheses("(())))))"))

    def test_large_inputs(self):
        large_balanced = "(" + "((((()()(((((())))))))))" * 100 + ")"
        large_unbalanced = "(()" + ")" * 100 + ")"
        self.assertTrue(balanced_parentheses(large_balanced))
        self.assertFalse(balanced_parentheses(large_unbalanced))


if __name__ == "__main__":
    unittest.main()
