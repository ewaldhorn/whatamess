import { initGame } from "./game";

// ------------------------------------------------------------------------------------------------
window.addEventListener("DOMContentLoaded", () => {
  console.log("Prelaunch message.");
  initGame();
  console.log("Should have rendered something.");
});
