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

console.log("Vowel count for 'hello' is expected to be 2:", getVowelCount("hello") === 2);