use bracket_lib::terminal::{BTerm, GameState};

#[derive(Debug)]
enum GameMode {
    Menu,
    Playing,
    End,
}

pub(crate) struct State {
    mode: GameMode,
}

impl State {
    pub(crate) fn new() -> Self {
        State {
            mode: GameMode::Menu,
        }
    }
    fn play(&mut self, ctx: &mut BTerm) {
        // TODO: Fill in this stub later
        self.mode = GameMode::End;
    }

    fn dead(&mut self, ctx: &mut BTerm) {
        // TODO: Fill in this stub later
        self.mode = GameMode::End;
    }

    fn main_menu(&mut self, ctx: &mut BTerm) {
        ctx.cls();
        ctx.print_centered(5, "Welcome to Crusty Rusty");
        ctx.print_centered(8, "(P) Play Game");
        ctx.print_centered(9, "(Q) Quit Game");
        self.mode = GameMode::End;
    }

    fn restart(&mut self) {
        // TODO: Fill in this stub later
        self.mode = GameMode::Playing;
    }
}

impl GameState for State {
    fn tick(&mut self, ctx: &mut BTerm) {
        match self.mode {
            GameMode::Menu => self.main_menu(ctx),
            GameMode::End => self.dead(ctx),
            GameMode::Playing => self.play(ctx),
        }
    }
}
