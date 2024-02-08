export const reverseString = (inp) => {
    // if the length is 0 or 1, return as-is
    if (inp.length < 2) {
        return inp;
    }

    // split string into characters
    // then reverse this character array
    // rejoin characters to form a string, now reversed
    return inp.split("").reverse().join("");
};
