from datetime import datetime


def display_current_time_raw():
    print(datetime.now())


def say_hello(who):
    print(f"Hello there {who}.")


display_current_time_raw()
say_hello("Bob")
