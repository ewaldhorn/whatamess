/**
 * A function that counts the vowels in a string.
 * 
 * getVowelCount("hello") => 2
 */

const getVowelCount = (string) => {
    const vowels = ['a', 'e', 'i', 'o', 'u'];
    let count = 0;
    for (let i = 0; i < string.length; i++) {
        if (vowels.includes(string[i].toLowerCase())) {
            count++;
        }
    }
    return count;
}

console.log("Vowel count for 'hello' is expected to be 2        :", getVowelCount("hello") === 2);
console.log("Vowel count for 'world' is expected to be 1        :", getVowelCount("world") === 1);
console.log("Vowel count for 'AEIOU' is expected to be 5        :", getVowelCount("AEIOU") === 5);
console.log("Vowel count for 'bcdfg' is expected to be 0        :", getVowelCount("bcdfg") === 0);
console.log("Vowel count for 'rhythm' is expected to be 0       :", getVowelCount("rhythm") === 0);
console.log("Vowel count for 'aEiOu' is expected to be 5        :", getVowelCount("aEiOu") === 5);
console.log("Vowel count for an empty string is expected to be 0:", getVowelCount("") === 0);

