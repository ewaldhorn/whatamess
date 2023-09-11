#![warn(clippy::all, clippy::pedantic)]
mod game;

use bracket_lib::prelude::*;
use game::State;

fn main() -> BError {
    let context = BTermBuilder::simple80x50().with_title("Crusty Rusty").build()?;
    main_loop(context, State{})
}
