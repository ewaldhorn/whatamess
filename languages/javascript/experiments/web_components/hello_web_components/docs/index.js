class b extends HTMLElement{constructor(){super();this.attachShadow({mode:"open"})}connectedCallback(){this.render()}render(){this.shadowRoot.innerHTML=`
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
        `}}class c extends HTMLElement{constructor(){super();this.attachShadow({mode:"open"})}connectedCallback(){this.render()}render(){this.shadowRoot.innerHTML=`
        <style>
            header {
              display: block;
              background-color: #333;
              color: #fff;
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
        `}}var i=()=>{customElements.define("nofuss-header",c),customElements.define("nofuss-footer",b)};i();
