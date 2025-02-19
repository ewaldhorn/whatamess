import { Engine } from "https://floaty.dev/engine-v1.js";
import { Points } from "./points";

// ----------------------------------------------------------------------------
//                                                            START YOUR ENGINE
let engine = new Engine();
engine.expose();

// ----------------------------------------------------------------------------
//                                                           SPRITES AND SOUNDS
let sprites = {};
let sounds = {};
/*
  This little Floaty gamelet doesn't use sounds or sprites, so empty objects
  are fine.
*/

// ----------------------------------------------------------------------------
//                                                                      GLOBALS
let mustClearScreen = true;
let points = new Points();

// ----------------------------------------------------------------------------
// //                                                          GET THINGS READY
async function init() {
  points.init();
  await this.load();
}

// ----------------------------------------------------------------------------
//                                                                 UPDATE STATE
function update() {
  points.update();
}

// ----------------------------------------------------------------------------
//                                                                 RENDER STATE
function draw() {
  if (mustClearScreen) {
    cls();
  }

  points.draw();

  // Clear the screen
  if (btnp("c")) {
    cls();
  }

  // Toggle overwrite or not
  if (btnp("o")) {
    mustClearScreen = !mustClearScreen;
  }
}

// ----------------------------------------------------------------------------
//                                                                  ENTRY POINT
engine.start({
  sprites,
  sounds,
  init,
  update,
  draw,
  target: document.querySelector(".game"),
});
