import os
import sys
import time
from datetime import datetime
from typing import Dict, List

class SystemInformationCollector:
    def __init__(self):
        self.system_info = {}

    def collect_system_info(self) -> None:
        self.system_info['os'] = os.name
        self.system_info['python_version'] = sys.version
        self.system_info['current_time'] = datetime.now().strftime('%Y-%m-%d %H:%M:%S')

    def get_system_info(self) -> Dict[str, str]:
        return self.system_info

class NameValidator:
    def __init__(self):
        self.validation_rules = [
            self.validate_length,
            self.validate_characters
        ]

    def validate_length(self, name: str) -> bool:
        return len(name) > 0

    def validate_characters(self, name: str) -> bool:
        return all(char.isalpha() or char.isspace() for char in name)

    def validate_name(self, name: str) -> bool:
        for rule in self.validation_rules:
            if not rule(name):
                return False
        return True

class GreetingGenerator:
    def __init__(self):
        self.greeting_templates = {
            'morning': 'Good morning',
            'afternoon': 'Good afternoon',
            'evening': 'Good evening'
        }

    def generate_greeting(self, name: str, current_time: datetime) -> str:
        hour = current_time.hour
        if hour < 12:
            return f"{self.greeting_templates['morning']}, {name}!"
        elif hour < 18:
            return f"{self.greeting_templates['afternoon']}, {name}!"
        else:
            return f"{self.greeting_templates['evening']}, {name}!"

class UserInputCollector:
    def __init__(self):
        self.name = None

    def collect_name(self) -> None:
        self.name = input("Please enter your name: ")

    def get_name(self) -> str:
        return self.name

class GreetingPrinter:
    def __init__(self):
        pass

    def print_greeting(self, greeting: str) -> None:
        print(greeting)

class ProgramRunner:
    def __init__(self):
        self.system_info_collector = SystemInformationCollector()
        self.name_validator = NameValidator()
        self.greeting_generator = GreetingGenerator()
        self.user_input_collector = UserInputCollector()
        self.greeting_printer = GreetingPrinter()

    def run(self) -> None:
        self.system_info_collector.collect_system_info()
        system_info = self.system_info_collector.get_system_info()
        print("System Information:")
        for key, value in system_info.items():
            print(f"{key.capitalize()}: {value}")
        time.sleep(2)
        self.user_input_collector.collect_name()
        name = self.user_input_collector.get_name()
        while not self.name_validator.validate_name(name):
            print("Invalid name. Please try again.")
            self.user_input_collector.collect_name()
            name = self.user_input_collector.get_name()
        current_time = datetime.now()
        greeting = self.greeting_generator.generate_greeting(name, current_time)
        print("\nGreeting:")
        self.greeting_printer.print_greeting(greeting)

def main():
    program_runner = ProgramRunner()
    program_runner.run()

if __name__ == "__main__":
    main()
