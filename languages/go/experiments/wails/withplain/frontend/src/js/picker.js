customElements.define(
  "random-picker",
  class extends HTMLElement {
    constructor() {
      super();

      this.uuid = `pick-${crypto.randomUUID()}`;

      this.innerHTML = `
        <form>
          <label for="${this.uuid}">Add a item</label>
          <input type="text" id="${this.uuid}">
          <button>Add Item</button>
        </form>
        <ul></ul>`;

      // get a handle on the elements
      this.form = this.querySelector("form");
      this.list = this.querySelector("ul");
    }
  },
);
