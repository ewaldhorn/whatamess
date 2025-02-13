import { NoFussFooter } from "./nofuss_footer.js";

// ----------------------------------------------------------------------------
const loadCustomElements = () => {
  customElements.define("nofuss-footer", NoFussFooter);
};

// ----------------------------------------------------------------------------
loadCustomElements();
