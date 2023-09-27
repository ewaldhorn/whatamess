/// <reference path="./global.d.ts" />
// @ts-check

export function cookingStatus(remainingTime) {
  if (remainingTime != undefined) {
    if (remainingTime == 0) {
      return "Lasagna is done.";
    } else {
      return "Not done, please wait.";
    }
  } else {
    return "You forgot to set the timer.";
  }
};

export function preparationTime(layers, prepTime) {
  if (prepTime === undefined) { prepTime = 2;};
  return layers.length * prepTime;  
}

export function quantities(layers) {
  let noodleCount = 0;
  let sauceCount = 0;

  layers.forEach((q) => {
    if (q === "noodles") {
      noodleCount += 1;
    } else if (q === "sauce") {
      sauceCount += 1;
    }
  });
  
  return {noodles: noodleCount * 50, sauce: sauceCount * 0.2};
}

export function addSecretIngredient(friendsList, myList) {
  myList.push(friendsList[friendsList.length - 1]);
}

export function scaleRecipe(recipe, portions) {
  let result = {};

  for (let key in recipe) {
    result[key] = (recipe[key] / 2.0) * portions;
  }
  
  return result;
}
