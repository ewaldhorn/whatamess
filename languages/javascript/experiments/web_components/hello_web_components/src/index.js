import { NoFussFooter } from "./nofuss_footer.js";
import { NoFussHeader } from "./nofuss_header.js";
import { NoFussWarningLight } from "./nofuss_warning_light.js";

// ----------------------------------------------------------------------------
const loadCustomElements = () => {
  customElements.define("nofuss-header", NoFussHeader);
  customElements.define("nofuss-footer", NoFussFooter);
  customElements.define("nofuss-warning-light", NoFussWarningLight);
};

// ----------------------------------------------------------------------------
loadCustomElements();
