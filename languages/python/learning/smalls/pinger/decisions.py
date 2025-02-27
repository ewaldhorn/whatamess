nopes = ["tony", "giselle"]


# -----------------------------------------------------------------------------
def grant_access(name: str, age: int, has_id: bool) -> bool:
    if not has_id:
        return False

    if name.lower() in nopes:
        print(f"You've been banned {name}!")
        return False

    if age >= 18:
        print(f"Welcome {name}!")
        return True

    return False


# -----------------------------------------------------------------------------
def main() -> None:
    assert grant_access("Niel", 19, False) is False
    assert grant_access("Daniel", 17, True) is False
    assert grant_access("Michelle", 28, False) is False
    assert grant_access("Giselle", 28, True) is False
    assert grant_access("Tony", 25, True) is False
    assert grant_access("Sandra", 19, True) is True
    assert grant_access("Dirko", 33, True) is True


# ================================================================= ENTRY POINT
if __name__ == "__main__":
    main()
