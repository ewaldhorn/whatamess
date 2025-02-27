from compounder import calculate_compound_interest


# -----------------------------------------------------------------------------
def get_starting_value() -> float:
    """Returns starting value from user.

    Returns:
        float: Starting value provided by user.
    """
    while True:
        try:
            num = float(input("Please enter a starting value: "))
            break
        except ValueError:
            print("That's not a valid number. Try again!")

    return num


# ================================================================= ENTRY POINT
if __name__ == "__main__":
    starting_value = get_starting_value()
    rate = 0.05
    years = 1
    print(starting_value)
    print(calculate_compound_interest(starting_value, rate, years))
