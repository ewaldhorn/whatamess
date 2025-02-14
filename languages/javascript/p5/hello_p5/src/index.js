import p5 from './p5.js';

// ----------------------------------------------------------------------------
const sketch = (p) => {

    let ball;
    let grid;

    // ------------------------------------------------------------------------ 
    p.setup = () => {
        p.createCanvas(400, 400);
        p.background(135, 206, 235); // Light blue background
        grid = [];
        for (let i = 0; i < 10; i++) {
            grid[i] = [];
            for (let j = 0; j < 10; j++) {
                grid[i][j] = 0;
            }
        }
        ball = {
            x: p.width / 2,
            y: p.height / 2,
            vx: 4,
            vy: 4,
            r: 20,
            c: p.color(255, 0, 0) // Red color
        };
    }

    // ------------------------------------------------------------------------
    p.draw = () => {
        p.background(135, 206, 235); // Light blue background
        // Draw grid
        for (let i = 0; i < 10; i++) {
            for (let j = 0; j < 10; j++) {
                p.noFill();
                p.stroke(255);
                p.rect(i * 40, j * 40, 40, 40);
            }
        }
        // Update and draw ball
        ball.x += ball.vx;
        ball.y += ball.vy;
        if (ball.x + ball.r > p.width || ball.x - ball.r < 0) {
            ball.vx *= -1;
        }
        if (ball.y + ball.r > p.height || ball.y - ball.r < 0) {
            ball.vy *= -1;
        }
        p.fill(ball.c);
        p.ellipse(ball.x, ball.y, ball.r * 2);
    }
}

// ----------------------------------------------------------------------------
new p5(sketch);