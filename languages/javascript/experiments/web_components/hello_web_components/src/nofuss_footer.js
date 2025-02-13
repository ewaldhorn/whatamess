export class NoFussFooter extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    const button = document.createElement("button");
    button.textContent = "Click me!";
    this.shadowRoot.appendChild(button);
  }
}
