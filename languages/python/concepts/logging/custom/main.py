import logging

# Log to a file and the console all at once.
# Handy for development purposes.

# Create a logger
logger = logging.getLogger(__name__)

# Set the logging level
logger.setLevel(logging.DEBUG)

# Create a file handler
file_handler = logging.FileHandler("delete_me.log")

# Create a console handler
console_handler = logging.StreamHandler()

# Create a formatter
formatter = logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - [%(message)s]")

# Add the formatter to the handlers
file_handler.setFormatter(formatter)
console_handler.setFormatter(formatter)

# Add the handlers to the logger
logger.addHandler(file_handler)
logger.addHandler(console_handler)

# Log some messages
logger.debug("This is a debug message")
logger.fatal("This is a disaster!")
