import os
import sys
import time
from datetime import datetime
from typing import Dict, List

class GreetingGenerator:
    def __init__(self):
        self.name = None
        self.greeting = None
        self.current_time = datetime.now()
        self.system_info = self.get_system_info()

    def get_system_info(self) -> Dict[str, str]:
        return {
            'os': os.name,
            'python_version': sys.version,
            'current_time': self.current_time.strftime('%Y-%m-%d %H:%M:%S')
        }

    def get_user_name(self) -> str:
        return input("Please enter your name: ")

    def validate_name(self, name: str) -> bool:
        return len(name) > 0 and all(char.isalpha() or char.isspace() for char in name)

    def generate_greeting(self, name: str) -> str:
        greetings = {
            'morning': 'Good morning',
            'afternoon': 'Good afternoon',
            'evening': 'Good evening'
        }
        hour = self.current_time.hour
        if hour < 12:
            return f"{greetings['morning']}, {name}!"
        elif hour < 18:
            return f"{greetings['afternoon']}, {name}!"
        else:
            return f"{greetings['evening']}, {name}!"

    def print_greeting(self):
        print(self.greeting)

    def run(self):
        self.name = self.get_user_name()
        while not self.validate_name(self.name):
            print("Invalid name. Please try again.")
            self.name = self.get_user_name()
        self.greeting = self.generate_greeting(self.name)
        print("System Information:")
        for key, value in self.system_info.items():
            print(f"{key.capitalize()}: {value}")
        time.sleep(2)
        print("\nGreeting:")
        self.print_greeting()

def main():
    generator = GreetingGenerator()
    generator.run()

if __name__ == "__main__":
    main()
