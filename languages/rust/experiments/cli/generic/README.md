# Generic project structure and options

Just a generic template I often use for new Rust projects.

## Cargo Watch

Command: `cargo add cargo-watch`

Now you can use `cargo watch -q -c -w src/ -x 'run -q'` to hot-reload changes.

Parameters are:

- q for quiet
- c clears screen
- w which files/folders to watch
- x which cargo command to execute on changes

## Useful crates

- serde serialize/deserialize
- serde_json specific to json
- thiserror adds a derive(Error) feature
- anyhow idiomatic error handling
- rayon makes parallel work easier
- tokio async framework for non-blocking IO

See <https://blessed.rs> for some possible Rust crates for new projects.
