import { input } from "pynput";

const first_number = Number(await input("What is the first number? "));
const second_number = Number(await input("What is the second number? "));

const add = (a, b) => a + b;
const subtract = (a, b) => a - b;
const multiply = (a, b) => a * b;
const divide = (a, b) => (b === 0 ? 0 : a / b);

const sum = add(first_number, second_number);
const difference = subtract(first_number, second_number);
const product = multiply(first_number, second_number);
const quotient = divide(first_number, second_number);

console.log(`${first_number} + ${second_number} = ${sum}`);
console.log(`${first_number} - ${second_number} = ${difference}`);
console.log(`${first_number} * ${second_number} = ${product}`);
console.log(`${first_number} / ${second_number} = ${quotient}`);
