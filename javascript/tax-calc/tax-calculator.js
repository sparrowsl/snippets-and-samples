import { input } from "pynput";

const tax = 5.5;

const amount = Number(await input("What is the order amount? "));
const state = await input("What is the state? ");

// Check the amount and state condition
if (state.toUpperCase() === "WI") {
	console.log(`The subtotal is $${amount.toFixed(2)}`);
	console.log(`The tax is $${((amount * tax) / 100).toFixed(2)}`);
	console.log(`The total is $${(amount + (amount * tax) / 100).toFixed(2)}`);
	process.exit();
}

console.log(`The total is $${amount.toFixed(2)}`);
