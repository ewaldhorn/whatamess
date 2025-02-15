  class MyComponent extends HTMLElement {
    constructor() {
      super();
      this.attachShadow({ mode: 'open' });
      this.shadowRoot.innerHTML = `
        <style>
          @import url('my-component.css');
        </style>
        <div class="my-component">
          <h1>Hello, World!</h1>
        </div>
      `;
    }
  }

  window.customElements.define('my-component', MyComponent);
