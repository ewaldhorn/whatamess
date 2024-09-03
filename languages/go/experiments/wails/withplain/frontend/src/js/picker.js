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
      this.notification = this.querySelector("#notification");
      this.field = this.form.querySelector("input");
      this.field.focus();

      // set up event listener(s)
      this.form.addEventListener("submit", this);
    }

    showNotification(msg) {
      let notification = document.createElement("div");
      notification.setAttribute("aria-live", "polite");
      notification.setAttribute("role", "status");

      this.form.append(notification);

      setTimeout(function () {
        notification.textContent = msg;
      }, 10);

      setTimeout(function () {
        notification.remove();
      }, 3000);
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

      this.showNotification(`Added "${this.field.value}"`);
      this.field.value = "";
    }
  },
);
