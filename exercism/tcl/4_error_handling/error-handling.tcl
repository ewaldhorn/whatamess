proc custom_error_message {message} {
    throw error $message
}
proc handle_error {script} {
    try {uplevel 1 $script} on ok {} {
        return "success"
    } trap {ARITH DIVZERO} {} {
        return "division by zero"
    } trap {POSIX ENOENT} {} {
        return "file does not exist"
    } trap {TCL LOOKUP COMMAND} {} {
        return "proc does not exist"
    }
}
