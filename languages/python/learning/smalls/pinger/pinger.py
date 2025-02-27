import time


# -----------------------------------------------------------------------------
def connect() -> None:
    print("Dial-up connecting to the internet")
    time.sleep(3)
    print("Connected!")


# ======================================================================= TESTS
if __name__ == "__main__":
    connect()
