from datetime import datetime


# -----------------------------------------------------------------------------
def display_current_time_raw():
    print(datetime.now())


# -----------------------------------------------------------------------------
def say_hello(who):
    print(f"Hello there {who}.")


# -----------------------------------------------------------------------------
def get_time_str():
    current_hour = datetime.now().hour

    match current_hour:
        case value if 6 <= value < 12:
            return "Good Morning"
        case value if 12 <= value < 18:
            return "Good Afternoon"
        case _:
            return "Good Evening"


# -----------------------------------------------------------------------------
def say_hello_time_sensitive(who):
    time_str = get_time_str()
    print(f"{time_str} {who}!")


# -----------------------------------------------------------------------------
display_current_time_raw()
say_hello("Bob")
say_hello_time_sensitive("Tina")
