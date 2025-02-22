import logging

# Create a logger
logger = logging.getLogger(__name__)

# Set the overall logging level to DEBUG
logger.setLevel(logging.DEBUG)

# Create a console handler
console_handler = logging.StreamHandler()
console_handler.setLevel(logging.DEBUG)

# Create a critical file handler
critical_handler = logging.FileHandler("critical.log")
critical_handler.setLevel(logging.CRITICAL)

# Create a general file handler
general_handler = logging.FileHandler("general.log")
general_handler.setLevel(logging.DEBUG)

# Create a formatter
formatter = logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s")

# Add the formatter to the handlers
console_handler.setFormatter(formatter)
critical_handler.setFormatter(formatter)
general_handler.setFormatter(formatter)

# Add the handlers to the logger
logger.addHandler(console_handler)
logger.addHandler(critical_handler)
logger.addHandler(general_handler)

# Log some messages
logger.debug("This is a debug message")
logger.info("This is an info message")
logger.warning("This is a warning message")
logger.error("This is an error message")
logger.critical("This is a critical message")
