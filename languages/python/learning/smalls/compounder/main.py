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
    print(starting_value)
