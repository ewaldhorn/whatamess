import logging

# ensure everything is logged, not just warnings and above, as per default
# also log to a file, because, well, yeah, that's always handy!
logging.basicConfig(filename='./delete_me.log', level=logging.DEBUG, format='%(asctime)s - %(levelname)s - %(message)s')

logging.debug("This is a debug message")
logging.info("This is an info message")
logging.warning("This is a warning message")
logging.error("This is an error message")
logging.critical("This is a critical message")
