customElements.define(
  "random-picker",
  class extends HTMLElement {
    // ------------------------------------------------------------------------
    constructor() {
      super();

      this.uuid = `pick-${crypto.randomUUID()}`;

      this.innerHTML = `
        <form>
          <label for="${this.uuid}">Add a item</label>
          <input type="text" id="${this.uuid}">
          <button>Add Item</button>
        </form>
        <ul></ul>
        <p><button pick-item>Pick Random Item</button></p>
        <div aria-live="polite" pick-result></div>`;

      // get a handle on the elements
      this.form = this.querySelector("form");
      this.list = this.querySelector("ul");
      this.notification = this.querySelector("#notification");
      this.field = this.form.querySelector("input");
      this.field.focus();
      this.pickButton = this.querySelector("[pick-item]");
      this.results = this.querySelector("[pick-result]");

      // set up event listener(s)
      this.form.addEventListener("submit", this);
      this.pickButton.addEventListener("click", this);
    }

    // ------------------------------------------------------------------------
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

    // ------------------------------------------------------------------------
    onsubmit(event) {
      event.preventDefault(); // prevents page reload

      if (this.field.value.trim().length > 0) {
        let newLi = document.createElement("li");
        newLi.textContent = this.field.value;
        this.list.append(newLi);

        this.showNotification(`Added "${this.field.value}"`);
        this.field.value = "";
      }
    }

    // ------------------------------------------------------------------------
    onclick(event) {
      this.results.innerHTML = "";
      let items = Array.from(this.list.querySelectorAll("li")).map(
        (item) => item.textContent,
      );
      this.results.innerHTML = `Random element: ${this.shuffleArray(items)[0]}`;
    }

    // ------------------------------------------------------------------------
    /**
     * Handle events
     * @param {Event} event The event object
     */
    handleEvent(event) {
      this[`on${event.type}`](event);
    }

    // ------------------------------------------------------------------------
    // Shuffles an array randomly
    shuffleArray(arr) {
      let currentIndex = arr.length;
      let tmpVal, randomIndex;

      while (0 != currentIndex) {
        // pick a random index
        randomIndex = Math.floor(Math.random() * currentIndex);
        currentIndex -= 1;

        // swap with current index
        tmpVal = arr[currentIndex];
        arr[currentIndex] = arr[randomIndex];
        arr[randomIndex] = tmpVal;
      }

      return arr;
    }
  },
);
