export class NoFussFooter extends HTMLElement {
  // --------------------------------------------------------------------------
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
  }

  // --------------------------------------------------------------------------
  connectedCallback() {
    this.render();
  }

  // --------------------------------------------------------------------------
  render() {
    const footerTemplate = `
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
                <li><a href="https://nofuss.co.za/" target="_blank">Home</a></li>
                <li><a href="https://nofuss.co.za/about/" target="_blank">About</a></li>
                <li><a href="https://nofuss.co.za/contact/" target="_blank">Contact</a></li>
              </ul>
            </nav>
          </footer>
        `;

    this.shadowRoot.innerHTML = footerTemplate;
  }
}
