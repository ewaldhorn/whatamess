def calculate_compound_interest(principal, rate, time, n=1):
    """
    Calculate compound interest.

    Args:
        principal (float): Initial amount
        rate (float): Interest rate (in decimal form, e.g., 4% = 0.04)
        time (int): Time period (in years)
        n (int, optional): Number of times interest is compounded per year. Defaults to 1.

    Returns:
        float: Final amount after compound interest
    """
    return principal * (1 + rate / n) ** (n * time)


# ======================================================================= TESTS

if __name__ == "__main__":
    principal = 10000
    rate = 0.04
    time = 5
    n = 12  # Monthly compounding
    expected = 12209.9659

    final_amount = calculate_compound_interest(principal, rate, time, n)
    rounded_amount = round(final_amount, 4)
    assert rounded_amount == expected
