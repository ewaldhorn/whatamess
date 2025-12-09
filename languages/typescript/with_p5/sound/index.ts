import p5 from "p5";

const sketch = function (p: p5) {
  p.setup = function () {
    p.createCanvas(p.windowWidth * 0.9, p.windowHeight * 0.9);
    p.background("black");
    p.noLoop();
  };

  p.draw = function () {
    p.strokeWeight(5);
    p.fill("red");
    p.rect(p.width / 2, p.height / 2, 100, 100);
  };
};

new p5(sketch);
