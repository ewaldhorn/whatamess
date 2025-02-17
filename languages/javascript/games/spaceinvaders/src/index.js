import { clearGameArea, clearScreen, drawCanvasBorder } from "./canvasutils.js";

// ----------------------------------------------------------------------------
//                                                                      GLOBALS
const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");

// ----------------------------------------------------------------------------
//                                                             CONFIGURE CANVAS
canvas.width = innerWidth - 50;
canvas.height = innerHeight - 100;

// ----------------------------------------------------------------------------
clearScreen(ctx);
drawCanvasBorder(ctx);
clearGameArea(ctx);
