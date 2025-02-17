class y extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: "open" }),
            this.currentVal = 0,
            this.maxVal = 100
    }
    static get observedAttributes() { return ["currentval", "maxval"] }
    attributeChangedCallback(o, f, l) {
        if (o === "currentval") this.currentVal = parseFloat(l);
        else if (o === "maxval") this.maxVal = parseFloat(l);
        this.render()
    }
    render() {
        let o = document.createElement("canvas");
        o.width = 35, o.height = 100;
        let f = o.getContext("2d");
        f.fillStyle = "rgba(99,99,99,1)",
            f.fillRect(0, 0, o.width, o.height),
            f.fillStyle = this.getFillColour(),
            f.fillRect(5, o.height - this.currentVal / this.maxVal * o.height, 25, this.currentVal / this.maxVal * o.height),
            this.shadowRoot.innerHTML = "",
            this.shadowRoot.appendChild(o)
    } getFillColour() { let o = this.currentVal / this.maxVal * 100; if (o < 50) return "rgba(255, 0, 0, 0.5)"; else if (o < 75) return "rgba(255, 255, 0, 0.5)"; else return "rgba(0, 255, 0, 0.5)" }
} 
    customElements.define("nofuss-warning-light", y) 