import p5 from './p5.js';

// ----------------------------------------------------------------------------
const spices = [
    "Basil",
    "Cinnamon",
    "Ginger",
    "Turmeric",
    "Paprika",
    "Rosemary",
    "Thyme",
    "Oregano",
    "Cumin",
    "Coriander",
    "Nutmeg",
    "Cardamom",
    "Cayenne Pepper",
    "Black Pepper",
    "White Pepper",
    "Garlic Powder",
    "Onion Powder",
    "Cloves",
    "Star Anise",
    "Fennel Seeds",
    "Mustard Seeds",
    "Dill Weed",
    "Saffron"
  ];

const sketch = (p) => {
    let balls = [];

    p.setup = () => {
        p.createCanvas(800, 800);
        p.background(135, 206, 235); // Light blue background

        // Generate a random number of balls 
        let numBalls = p.floor(p.random(8, 21));

        // Create the balls
        balls = [];
        for (let i = 0; i < numBalls; i++) {
            // Generate random size, color, and initial position
            let r = p.random(10, 70); // Radius between 10 and 70
            let c = p.color(p.random(255), p.random(255), p.random(255)); // Random color
            let x = p.random(r, p.width - r); // Initial x position
            let y = p.random(r, p.height - r); // Initial y position
            let vx = p.random(-4, 4); // Initial x velocity
            let vy = p.random(-4, 4); // Initial y velocity
            let name = `${spices[i]}`; // Spice name

            // Add the ball to the array
            balls.push({
                x: x,
                y: y,
                vx: vx,
                vy: vy,
                r: r,
                c: c,
                name: name,
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
                p.stroke(115, 186, 215);
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

            // Draw label
            p.fill(0); // Black text color
            p.textAlign(p.CENTER, p.CENTER);
            p.text(balls[i].name, balls[i].x, balls[i].y + balls[i].r + 10);
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
new p5(sketch, "p5-main-canvas");