import { darkTheme } from "./theme";

export class NoFussWarningLight extends HTMLElement {
  // --------------------------------------------------------------------------
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.currentVal = 0;
    this.maxVal = 100;
  }

  // --------------------------------------------------------------------------
  static get observedAttributes() {
    return ["currentval", "maxval"];
  }

  // --------------------------------------------------------------------------
  attributeChangedCallback(name, oldValue, newValue) {
    if (name === "currentval") {
      this.currentVal = parseFloat(newValue);
    } else if (name === "maxval") {
      this.maxVal = parseFloat(newValue);
    }
    this.render();
  }

  // --------------------------------------------------------------------------
  render() {
    const canvas = document.createElement("canvas");
    canvas.width = 35;
    canvas.height = 100;
    const ctx = canvas.getContext("2d");

    // Clear the canvas
    ctx.fillStyle = "rgba(99,99,99,1)";
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    // Draw the graph
    ctx.fillStyle = this.getFillColour();
    ctx.fillRect(
      5,
      canvas.height - (this.currentVal / this.maxVal) * canvas.height,
      25,
      (this.currentVal / this.maxVal) * canvas.height,
    );

    // Update the shadow DOM
    this.shadowRoot.innerHTML = "";
    this.shadowRoot.appendChild(canvas);
  }

  // --------------------------------------------------------------------------
  getFillColour() {
    const percentage = (this.currentVal / this.maxVal) * 100;

    if (percentage < 50) {
      return "rgba(255, 0, 0, 0.5)";
    } else if (percentage < 75) {
      return "rgba(255, 255, 0, 0.5)";
    } else {
      return "rgba(0, 255, 0, 0.5)";
    }
  }
}
