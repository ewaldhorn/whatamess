import { NoFussFooter } from "./nofuss_footer.js";
import { NoFussHeader } from "./nofuss_header.js";

// ----------------------------------------------------------------------------
const loadCustomElements = () => {
  customElements.define("nofuss-header", NoFussHeader);
  customElements.define("nofuss-footer", NoFussFooter);
};

// ----------------------------------------------------------------------------
loadCustomElements();
