// the "n" at the end means it's a BigInt
const bigInt = 1234567890123456789012345678901234567890n;

const sameBigint = BigInt("1234567890123456789012345678901234567890");

const bigintFromNumber = BigInt(10); // same as 10n

alert(1n + 2n); // 3

alert(5n / 2n); // 2

alert(1n + 2); // Error: Cannot mix BigInt and other types

let bigint = 1n;
let number = 2;

// number to bigint
alert(bigint + BigInt(number)); // 3

// bigint to number
alert(Number(bigint) + number); // 3

let bigint = 1n;

alert( +bigint ); // error

alert( 2n > 1n ); // true

alert( 2n > 1 ); // true

alert( 1 == 1n ); // true

alert( 1 === 1n ); // false

if (0n) {
  // never executes
}

alert( 1n || 2 ); // 1 (1n is considered truthy)

alert( 0n || 2 ); // 2 (0n is considered falsy)


