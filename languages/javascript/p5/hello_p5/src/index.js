import p5 from './p5.js';

// ----------------------------------------------------------------------------
const sketch = (p) => {
    let balls = [];

    p.setup = () => {
        p.createCanvas(600, 600);
        p.background(135, 206, 235); // Light blue background

        // Generate a random number of balls between 5 and 8
        let numBalls = p.floor(p.random(5, 9));

        // Create the balls
        balls = [];
        for (let i = 0; i < numBalls; i++) {
            // Generate random size, color, and initial position
            let r = p.random(10, 30); // Radius between 10 and 30
            let c = p.color(p.random(255), p.random(255), p.random(255)); // Random color
            let x = p.random(r, p.width - r); // Initial x position
            let y = p.random(r, p.height - r); // Initial y position
            let vx = p.random(-4, 4); // Initial x velocity
            let vy = p.random(-4, 4); // Initial y velocity

            // Add the ball to the array
            balls.push({
                x: x,
                y: y,
                vx: vx,
                vy: vy,
                r: r,
                c: c
            });
        }
    };

    p.draw = () => {
        p.background(135, 206, 235); // Light blue background

        // Draw grid
        let gridSize = 40; // Grid cell size
        let numCols = p.floor(p.width / gridSize);
        let numRows = p.floor(p.height / gridSize);

        for (let i = 0; i < numCols; i++) {
            for (let j = 0; j < numRows; j++) {
                p.noFill();
                p.stroke(255);
                p.rect(i * gridSize, j * gridSize, gridSize, gridSize);
            }
        }

        // Update and draw balls
        for (let i = 0; i < balls.length; i++) {
            balls[i].x += balls[i].vx;
            balls[i].y += balls[i].vy;

            // Bounce off edges
            if (balls[i].x + balls[i].r > p.width || balls[i].x - balls[i].r < 0) {
                balls[i].vx *= -1;
            }
            if (balls[i].y + balls[i].r > p.height || balls[i].y - balls[i].r < 0) {
                balls[i].vy *= -1;
            }

            // Draw ball
            p.fill(balls[i].c);
            p.ellipse(balls[i].x, balls[i].y, balls[i].r * 2);
        }

        // Check for collisions between balls
        for (let i = 0; i < balls.length; i++) {
            for (let j = i + 1; j < balls.length; j++) {
                let dx = balls[i].x - balls[j].x;
                let dy = balls[i].y - balls[j].y;
                let distance = p.sqrt(dx * dx + dy * dy);

                if (distance < balls[i].r + balls[j].r) {
                    // Swap velocities
                    let tempVx = balls[i].vx;
                    let tempVy = balls[i].vy;
                    balls[i].vx = balls[j].vx;
                    balls[i].vy = balls[j].vy;
                    balls[j].vx = tempVx;
                    balls[j].vy = tempVy;
                }
            }
        }
    };
};

// ----------------------------------------------------------------------------
new p5(sketch);