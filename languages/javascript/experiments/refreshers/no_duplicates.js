export const noDuplicates = (str) => {
    const charactersSeen = {};

    for (const ch of str) {
        if (ch in charactersSeen) {
            return false;
        }

        charactersSeen[ch] = true;
    }

    return true;
}

export const noDuplicatesAlternative = (str) => {
    for (let src = 0; src < str.length; src++) {
        for (let chk = src + 1; chk < str.length; chk++) {
            if (str[chk] == str[src]) {
                return false;
            }
        }
    }

    return true;
}

export const noDuplicatesSorted = (str) => {
    const tmp = str.split("").sort().join("");

    for (let pos = 0; pos < tmp.length; pos++) {
        if (tmp[pos] === tmp[pos + 1]) {
            return false;
        }
    }

    return true;
}