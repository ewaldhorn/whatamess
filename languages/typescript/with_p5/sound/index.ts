import p5 from "p5";

// 1. TypeScript FIX: Safely extract the instance type from the default export
type P5Instance = InstanceType<typeof p5>;

// Define the P5 instance type with the added sound properties for type safety.
// We assert that the p5 instance will include the sound library's components.
interface P5SoundInstance extends P5Instance {
  Oscillator: new (type?: string) => any; // Use 'any' or a more specific sound library type if you define one
  // Add other sound properties like AudioIn, Effect, etc., if you use them
}

// Create a function that defines your p5 sketch logic
const sketch = (p: P5SoundInstance) => {
  let oscillator: any; // Using 'any' here for simplicity with the p5.sound object

  p.setup = () => {
    p.createCanvas(400, 400);
    p.background(220);

    // p5.sound usage
    // We can safely use p.Oscillator because we asserted the type via P5SoundInstance
    oscillator = new p.Oscillator("sine");
    oscillator.amp(0);
    oscillator.start();

    p.text("Click to start sound", 10, 20);
  };

  p.draw = () => {
    p.fill(p.random(255), p.random(255), p.random(255));
    p.ellipse(p.mouseX, p.mouseY, 50, 50);
  };

  p.mousePressed = () => {
    if (oscillator) {
      oscillator.amp(0.5, 0.05); // Ramp up to 0.5 amplitude
    }
  };
};

// Initialize the p5 instance using the specialized sketch function
new p5(sketch as unknown as (p: p5) => void);
