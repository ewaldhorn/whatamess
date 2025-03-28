use bracket_lib::prelude::*;

pub const SCREEN_WIDTH: i32 = 80;
pub const SCREEN_HEIGHT: i32 = 50;
pub const FRAME_DURATION: f32 = 75.0;
const PLAYER_FRAMES: [u16; 6] = [1, 2, 3, 4, 2, 1];

struct Obstacle {
    x: i32,
    gap_y: i32,
    size: i32,
}

impl Obstacle {
    fn new(x: i32, score: i32) -> Self {
        let mut random = RandomNumberGenerator::new();
        Obstacle {
            x,
            gap_y: random.range(10, 40),
            size: i32::max(2, 20 - score),
        }
    }

    fn render(&mut self, ctx: &mut BTerm, player_x: i32) {
        let screen_x = self.x - player_x;
        let half_size = self.size / 2;

        // ground
        for x in 0..SCREEN_WIDTH {
            ctx.set(x, SCREEN_HEIGHT - 1, WHITE, WHITE, to_cp437('#'));
        }

        for y in 0..self.gap_y - half_size {
            // draw top
            ctx.set(screen_x, y, RED, BLACK, 179);
        }

        for y in self.gap_y + half_size..SCREEN_HEIGHT - 1 {
            // draw bottom
            ctx.set(screen_x, y, RED, BLACK, 179);
        }
    }

    fn hit_obstacle(&self, player: &Player) -> bool {
        let half_size = self.size / 2;
        let does_x_match = player.x == self.x;
        let player_above_gap = player.y < self.gap_y - half_size;
        let player_below_gap = player.y > self.gap_y + half_size;
        does_x_match && (player_above_gap || player_below_gap)
    }
}

struct Player {
    x: i32,
    y: i32,
    velocity: f32,
    frame: usize,
}

impl Player {
    fn new(x: i32, y: i32) -> Self {
        Player {
            x,
            y,
            velocity: 0.0,
            frame: 0,
        }
    }

    fn render(&mut self, ctx: &mut BTerm) {
        ctx.set(0, self.y, YELLOW, BLACK, PLAYER_FRAMES[self.frame]);
    }

    fn apply_gravity_and_move(&mut self) {
        if self.velocity < 2.0 {
            self.velocity += 0.2;
        }

        self.y += self.velocity as i32; // rounds down to an integer
        self.x += 1;

        if self.y < 0 {
            self.y = 0;
        }

        self.frame += 1;
        if self.frame > 5 {
            self.frame = 0;
        }
    }

    fn flap(&mut self) {
        self.velocity = -2.0;
    }
}

#[derive(Debug)]
enum GameMode {
    Menu,
    Playing,
    End,
}

pub(crate) struct State {
    player: Player,
    frame_time: f32,
    obstacle: Obstacle,
    mode: GameMode,
    score: i32,
}

impl State {
    pub(crate) fn new() -> Self {
        State {
            player: Player::new(5, 25),
            frame_time: 0.0,
            obstacle: Obstacle::new(SCREEN_WIDTH, 0),
            mode: GameMode::Menu,
            score: 0,
        }
    }
    fn play(&mut self, ctx: &mut BTerm) {
        ctx.cls_bg(BLACK);
        self.frame_time += ctx.frame_time_ms;

        if self.frame_time > FRAME_DURATION {
            self.frame_time = 0.0;
            self.player.apply_gravity_and_move();
        }

        if let Some(VirtualKeyCode::Space) = ctx.key {
            self.player.flap();
        }

        self.obstacle.render(ctx, self.player.x);
        self.player.render(ctx);
        ctx.print(0, 0, "Press <SPACE> to boost.");
        ctx.print(0, 1, &format!("Your score is {}", self.score));

        if self.player.x > self.obstacle.x {
            self.score += 1;
            self.obstacle = Obstacle::new(self.player.x + SCREEN_WIDTH, self.score);
        }

        if self.player.y > SCREEN_HEIGHT || self.obstacle.hit_obstacle(&self.player) {
            self.mode = GameMode::End;
        }
    }

    fn dead(&mut self, ctx: &mut BTerm) {
        ctx.cls();
        ctx.print_centered(5, "Oh no, you crashed!");
        ctx.print_centered(6, &format!("Your final score was {}", self.score));
    }

    fn main_menu(&mut self, ctx: &mut BTerm) {
        ctx.cls();
        ctx.print_centered(5, "Welcome to Crusty Rusty");
        ctx.print_centered(8, "(P) Play Game");
        ctx.print_centered(9, "(Q) Quit Game");

        if let Some(key) = ctx.key {
            match key {
                VirtualKeyCode::P => self.restart(),
                VirtualKeyCode::Q => ctx.quitting = true,
                _ => {}
            }
        }
    }

    fn restart(&mut self) {
        self.player = Player::new(5, 25);
        self.frame_time = 0.0;
        self.obstacle = Obstacle::new(SCREEN_WIDTH, 0);
        self.score = 0;
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
