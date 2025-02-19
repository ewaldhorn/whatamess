/*
"Old" way to "new" way.

Using JavaScript for functional programming.
*/
// ============================================================================
//
// Loops

// traditional
const numbers = [1, 2, 3, 4, 5];
const doubledNumbers = [];

for (let i = 0; i < numbers.length; i++) {
  doubledNumbers.push(numbers[i] * 2);
}

// functional
const numbers = [1, 2, 3, 4, 5];
const doubledNumbers = numbers.map(n => n * 2);
// ============================================================================
//
// Conditional statements

// traditional
function greet(name) {
  let message;
  if (name) {
    message = `Hello, ${name}!`;
  } else {
    message = 'Hello!';
  }
  return message;
}

// functional
function greet(name) {
  return name ? `Hello, ${name}!` : 'Hello!';
}

// ============================================================================
//
// Object property access

// traditional
const person = { name: 'John', age: 30 };
const name = person.name;
const age = person.age;

// functional
const person = { name: 'John', age: 30 };
const { name, age } = person;

// ============================================================================
//
// Array indexing

// traditional
const numbers = [1, 2, 3];
const first = numbers[0];
const rest = numbers.slice(1);

// functional
const numbers = [1, 2, 3];
const [first, ...rest] = numbers;

// ============================================================================
//
// State

// traditional - Mutable
const person = { name: 'John', age: 30 };
person.age += 1;

// functional - Immutable
const person = { name: 'John', age: 30 };
const updatedPerson = { ...person, age: person.age + 1 };

// ============================================================================
// ============================================================================
// ============================================================================
