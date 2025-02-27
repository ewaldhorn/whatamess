nopes = ["tony", "giselle"]


# -----------------------------------------------------------------------------
def has_id_and_is_old_enough(age: int, has_id: bool) -> bool:
    if has_id and age >= 18:
        return True
    else:
        return False


# -----------------------------------------------------------------------------
def is_banned(name: str) -> bool:
    return name.lower() in nopes


# -----------------------------------------------------------------------------
def grant_access(name: str, age: int, has_id: bool) -> bool:
    if not is_banned(name) and has_id_and_is_old_enough(age, has_id):
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
