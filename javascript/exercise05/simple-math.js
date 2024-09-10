import { input } from "pynput";

const firstNumber = Number(await input("What is the first number? "));
const secondNumber = Number(await input("What is the second number? "));

const sum = add(firstNumber, secondNumber);
const difference = subtract(firstNumber, secondNumber);
const product = multiply(firstNumber, secondNumber);
const quotient = divide(firstNumber, secondNumber);

console.log(`${firstNumber} + ${secondNumber} = ${sum}`);
console.log(`${firstNumber} - ${secondNumber} = ${difference}`);
console.log(`${firstNumber} * ${secondNumber} = ${product}`);
console.log(`${firstNumber} / ${secondNumber} = ${quotient}`);

function add(a, b) {
	return a + b;
}

function subtract(a, b) {
	return a - b;
}

function multiply(a, b) {
	return a * b;
}

function divide(a, b) {
	if (b === 0) {
		return 0;
	}
	return a / b;
}
