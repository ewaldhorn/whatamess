import globals from "./globals.ts";

// ------------------------------------------------------------------------------------------------
const setGameInformation = () => {
  document.title = `${globals.GAME_NAME} v${globals.VERSION}`;
};

// ------------------------------------------------------------------------------------------------
export const initGame = () => {
  setGameInformation();
};
