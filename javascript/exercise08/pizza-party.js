import { input } from "pynput";

const people = Number(await input("How many people? "));
if (!people) {
	console.log("Please enter a valid number!!");
	process.exit();
}

const pizzas = Number(await input("How many pizzas do you have? "));
if (!pizzas) {
	console.log("Please enter a valid number!!");
	process.exit();
}

const leftovers = people % pizzas;
console.log(
	`${people} people with ${pizzas} pizzas\nEach person gets ${people / pizzas / 2} pieces of pizza.`,
);
console.log(`There are ${leftovers} leftovers pieces.`);
