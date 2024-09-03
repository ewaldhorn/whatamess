import { greet } from "/js/greet.js";

document.addEventListener("DOMContentLoaded", (event) => {
  console.log("DOM fully loaded and parsed");

  document.getElementById("actionButton").onclick = greet;
});
