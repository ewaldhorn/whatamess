# Cargo Watch

Command: `cargo add cargo-watch`

Now you can use `cargo watch -q -c -w src/ -x 'run -q'` to hot-reload changes.

Parameters are:

- q for quiet
- c clears screen
- w which files/folders to watch
- x which cargo command to execute on changes
