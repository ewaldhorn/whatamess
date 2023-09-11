#![warn(clippy::all, clippy::pedantic)]

extern crate rand;

use rand::distributions::{IndependentSample, Range};
use rand::Rng;
use std::time::Instant;

struct Point {
    x: f64,
    y: f64,
}

impl Point {
    // our random point is just two random numbers uniformly chosen
    // between -1 and 1
    fn random<R: Rng>(rng: &mut R) -> Point {
        let distribution = Range::new(-1f64, 1.);

        Point {
            x: distribution.ind_sample(rng),
            y: distribution.ind_sample(rng),
        }
    }

    // Pythagoras' theorem at work
    fn is_in_unit_circle(&self) -> bool {
        self.x * self.x + self.y * self.y <= 1.
    }
}

fn calculate_pi<R: Rng>(rng: &mut R, total: u64) -> f64 {
    // this will generate the random points and count the number that
    // fall in the unit circle
    let in_circle = (0..total)
        .map(|_| Point::random(rng))
        .filter(|p|Point::is_in_unit_circle(p))
        .count();

    // knowing that, we can solve for pi
    let square_area = 4.;
    square_area * (in_circle as f64) / (total as f64)
}

fn main() {
    let mut rng = rand::thread_rng();

    // more iterations are more accurate, but will take longer to run
    for i in 0..10 {
        let base: u64 = 10;
        let iterations = base.pow(i);

        let start_time = Instant::now();
        let pi = calculate_pi(&mut rng, iterations);
        let duration = start_time.elapsed();

        println!(
            "Iterations: {:10}, Pi: {:10.8}, Time: {}s{:09}ns",
            iterations,
            pi,
            duration.as_secs(),
            duration.subsec_nanos()
        );
    }
}
