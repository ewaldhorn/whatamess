use bracket_lib::terminal::{GameState, BTerm};

pub(crate) struct State {}

impl GameState for State {
    fn tick(&mut self, ctx: &mut BTerm) {
        ctx.cls();
        ctx.print(1,1,"Howdy there Bracket Terminal!");
    }
}