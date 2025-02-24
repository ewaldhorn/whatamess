from datetime import datetime


def display_current_time_raw():
    print(datetime.now())


def say_hello(who):
    print(f"Hello there {who}.")


def get_time_str():
    return "Good morning"


def say_hello_time_sensitive(who):
    time_str = get_time_str()
    print(f"{time_str} {who}!")


display_current_time_raw()
say_hello("Bob")
say_hello_time_sensitive("Tina")
