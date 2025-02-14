import p5 from './p5.js';

// ----------------------------------------------------------------------------
let herbs = ["Garlic", "Thyme", "Rosemary", "Parsley"];
let data = [
    [1.00, 0.35, 0.25, 0.30],
    [0.35, 1.00, 0.20, 0.25],
    [0.25, 0.20, 1.00, 0.20],
    [0.30, 0.25, 0.20, 1.00]
];
const sketch = (p) => {
    p.setup = () => {
        p.createCanvas(400, 400);
        p.background(255);
        p.noStroke();
        drawHeatmap();
        p.filter(p.BLUR, 1);
    };

    function drawHeatmap() {
        let cellSize = p.width / herbs.length;

        for (let i = 0; i < herbs.length; i++) {
            for (let j = 0; j < herbs.length; j++) {
                let value = data[i][j];
                let colorValue = p.map(value, 0, 1, 0, 255);
                p.fill(colorValue, 0, 255 - colorValue);
                p.rect(j * cellSize, i * cellSize, cellSize, cellSize);
            }
        }

        // Add labels
        p.fill(255);
        p.textSize(16);
        p.noStroke();
        for (let i = 0; i < herbs.length; i++) {
            p.text(herbs[i], 10, i * cellSize + cellSize / 2 + 5);
            // p.textAlign(p.CENTER, p.BOTTOM);
            p.text(herbs[i], i * cellSize + cellSize / 2 - 50, p.height - 10);
        }
    }
};

// ---------------------------------------------------------------------------
const scatterPlot = (p) => {
     p.setup =()=> {
        p.createCanvas(800, 800);
        p.background(255);
        drawScatterPlotMatrix();
      }
      
      function drawScatterPlotMatrix() {
        let numHerbs = herbs.length;
        let cellSize = p.width / numHerbs;
        
        for (let i = 0; i < numHerbs; i++) {
          for (let j = 0; j < numHerbs; j++) {
            let x = j * cellSize;
            let y = i * cellSize;
            
            // Draw scatter plot
            p.fill(0);
            p.noStroke();
            for (let k = 0; k < 100; k++) {
              let vx = p.randomGaussian() * 0.1;
              let vy = p.randomGaussian() * 0.1;
              let correlation = data[i][j];
              vx *= correlation;
              vy *= correlation;
              p.ellipse(x + cellSize / 2 + vx * cellSize / 2, y + cellSize / 2 + vy * cellSize / 2, 2, 2);
            }
            
            // Add labels
            p.fill(0);
            p.textSize(12);
            p.textAlign(p.CENTER, p.TOP);
            p.text(herbs[j], x + cellSize / 2, y + 5);
            p.textAlign(p.LEFT, p.CENTER);
            p.text(herbs[i], x - 20, y + cellSize / 2);
          }
        }
      }
}

// ----------------------------------------------------------------------------
new p5(sketch, "p5-main-canvas");
new p5(scatterPlot, "p5-scatter-plot");