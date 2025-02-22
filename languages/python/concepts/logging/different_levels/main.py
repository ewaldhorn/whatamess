import logging
import os
from datetime import datetime

# set up a daily log file
# use different levels for files vs console
# console logs everything
# file logs only contains warnings or higher

def setup_daily_logger():
    base_dir = os.path.dirname(os.path.abspath(__file__))
    log_dir = os.path.join(base_dir, 'logs')
    os.makedirs(log_dir, exist_ok=True)

    current_time = datetime.now().strftime("%m_%d_%y_%I_%M_%p")
    log_file = os.path.join(log_dir, f"{current_time}.log")

    # Create the logger, set default log level here
    logger = logging.getLogger(__name__)
    logger.setLevel(logging.DEBUG)

    # Create a console handler
    console_handler = logging.StreamHandler()
    console_handler.setLevel(logging.DEBUG)  # Log all messages to the console

    # Create a file handler
    # all logs live in the 'logs' folder
    # log only warnings or higher to the file
    file_handler = logging.FileHandler(log_file)
    file_handler.setLevel(logging.WARNING)

    # Create a formatter
    formatter = logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s")

    # Add the formatter to the handlers
    console_handler.setFormatter(formatter)
    file_handler.setFormatter(formatter)

    # Add the handlers to the logger
    logger.addHandler(console_handler)
    logger.addHandler(file_handler)

    return logging.getLogger(__name__)

logger = setup_daily_logger()
logger.debug("Not for the file")
logger.info("Also not for the file")
logger.warning("And this goes to the file!")
