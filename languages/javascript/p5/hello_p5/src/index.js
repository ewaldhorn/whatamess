import p5 from './p5.js';

// ----------------------------------------------------------------------------
const sketch = (p) => {
    let balls = [];
  
    p.setup = () => {
      p.createCanvas(400, 400);
      p.background(135, 206, 235); // Light blue background
  
      // Create two balls
      balls.push({
        x: p.width / 2,
        y: p.height / 2,
        vx: 4,
        vy: 4,
        r: 20,
        c: p.color(255, 0, 0) // Red color
      });
  
      balls.push({
        x: p.width / 4,
        y: p.height / 4,
        vx: -3,
        vy: 5,
        r: 20,
        c: p.color(0, 0, 255) // Blue color
      });

      balls.push({
        x: p.width / 3,
        y: p.height / 4,
        vx: -2,
        vy: 2,
        r: 10,
        c: p.color(0, 255, 0) // Green color
      });           
    };
  
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