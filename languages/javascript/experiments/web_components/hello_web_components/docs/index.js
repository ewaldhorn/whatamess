class r extends HTMLElement{constructor(){super();this.attachShadow({mode:"open"})}connectedCallback(){this.render()}render(){this.shadowRoot.innerHTML=`
        <style>
            footer {
              display: block;
              background-color: #333;
              color: #fff;
              padding: 20px;
              text-align: center;
            }

            footer {
              max-width: 1000px;
              margin: 0 auto;
            }

            footer nav ul {
              list-style: none;
              margin: 0;
              padding: 0;
              display: flex;
              justify-content: space-around;
            }

            footer nav ul li {
              margin-right: 20px;
            }

            footer nav a {
              color: #fff;
              text-decoration: none;
            }
          </style>

          <footer>
            <p>&copy; 2025 NoFuss Solutions</p>
            <nav>
              <ul>
                <li><a href="https://nofuss.co.za/" target="_blank">Home</a></li>
                <li><a href="https://nofuss.co.za/about/" target="_blank">About</a></li>
                <li><a href="https://nofuss.co.za/contact/" target="_blank">Contact</a></li>
              </ul>
            </nav>
          </footer>
        `}}var C={primaryColour:"#2f2f2f",secondaryColour:"#4f4f4f",textColour:"#cccccc",accentColour:"#66d9ef",borderColour:"#333333",hoverColour:"#5a5a5a",activeColour:"#8f8f8f",disabledColour:"#777777"};class b extends HTMLElement{constructor(){super();this.attachShadow({mode:"open"})}connectedCallback(){this.render()}render(){let o=`
        <style>
            :host {
                background-color: ${C.primaryColour};
                color: ${C.textColour};
            }

            header {
              display: block;
              background-color: ${C.accentColor};
              color: ${C.textColour};
              padding: 20px;
              text-align: center;
            }

            header {
              max-width: 1000px;
              margin: 0 auto;
            }

            header nav ul {
              list-style: none;
              margin: 0;
              padding: 0;
              display: flex;
              justify-content: space-around;
            }

            header nav ul li {
              margin-right: 20px;
            }

            header nav a {
              color: #fff;
              text-decoration: none;
            }
          </style>

          <header>
            <h1>NoFuss Solutions</h1>
            <h2>Web Components Experiment</h2>
            <nav>
              <ul>
                <li><a href="https://nofuss.co.za/" target="_blank">Home</a></li>
                <li><a href="https://nofuss.co.za/about/" target="_blank">About</a></li>
                <li><a href="https://nofuss.co.za/contact/" target="_blank">Contact</a></li>
              </ul>
            </nav>
          </header>
        `;this.shadowRoot.innerHTML=o}}class y extends HTMLElement{constructor(){super();this.attachShadow({mode:"open"}),this.currentVal=0,this.maxVal=100}static get observedAttributes(){return["currentval","maxval"]}attributeChangedCallback(o,f,l){if(o==="currentval")this.currentVal=parseFloat(l);else if(o==="maxval")this.maxVal=parseFloat(l);this.render()}render(){let o=document.createElement("canvas");o.width=35,o.height=100;let f=o.getContext("2d");f.fillStyle="rgba(99,99,99,1)",f.fillRect(0,0,o.width,o.height),f.fillStyle=this.getFillColour(),f.fillRect(5,o.height-this.currentVal/this.maxVal*o.height,25,this.currentVal/this.maxVal*o.height),this.shadowRoot.innerHTML="",this.shadowRoot.appendChild(o)}getFillColour(){let o=this.currentVal/this.maxVal*100;if(o<50)return"rgba(255, 0, 0, 0.5)";else if(o<75)return"rgba(255, 255, 0, 0.5)";else return"rgba(0, 255, 0, 0.5)"}}var q=()=>{customElements.define("nofuss-header",b),customElements.define("nofuss-footer",r),customElements.define("nofuss-warning-light",y)};q();
