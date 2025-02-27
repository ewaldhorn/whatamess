# -----------------------------------------------------------------------------
def upper_case(items: list[str]) -> list[str]:
    return [item.upper() for item in items]


# ================================================================= ENTRY POINT
if __name__ == "__main__":
    print(upper_case(["ted", "bill", "10"]))
