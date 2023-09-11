#![warn(clippy::all, clippy::pedantic)]
mod game;

use bracket_lib::prelude::*;
use game::{State, SCREEN_HEIGHT, SCREEN_WIDTH};

fn main() -> BError {
    let context = BTermBuilder::new()
        .with_font("../resources/sprite_sheet_32.png", 32, 32)
        .with_simple_console(
            SCREEN_WIDTH,
            SCREEN_HEIGHT,
            "../resources/sprite_sheet_32.png",
        )
        .with_fancy_console(
            SCREEN_WIDTH,
            SCREEN_HEIGHT,
            "../resources/sprite_sheet_32.png",
        )
        .with_title("Crusty Rusty")
        .with_tile_dimensions(16, 16)
        .build()?;
    main_loop(context, State::new())
}
