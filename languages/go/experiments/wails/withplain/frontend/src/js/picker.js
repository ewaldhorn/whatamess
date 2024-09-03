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
      this.field = this.form.querySelector("input");
      this.field.focus();

      // set up event listener(s)
      this.form.addEventListener("submit", this);
    }

    /**
     * Handle events
     * @param {Event} event The event object
     */
    handleEvent(event) {
      event.preventDefault(); // prevents page reload

      let newLi = document.createElement("li");
      newLi.textContent = this.field.value;
      this.list.append(newLi);

      this.field.value = "";
    }
  },
);
