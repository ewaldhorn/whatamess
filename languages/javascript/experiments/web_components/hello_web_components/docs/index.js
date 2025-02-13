class b extends HTMLElement{constructor(){super();this.attachShadow({mode:"open"})}connectedCallback(){this.render()}render(){this.shadowRoot.innerHTML=`
        <style>
          /* NoFuss Footer Element */
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
                <li><a href="https://nofuss.co.za/">Home</a></li>
                <li><a href="https://nofuss.co.za/about/">About</a></li>
                <li><a href="https://nofuss.co.za/contact/">Contact</a></li>
              </ul>
            </nav>
          </footer>
        `}}var c=()=>{customElements.define("nofuss-footer",b)};c();
