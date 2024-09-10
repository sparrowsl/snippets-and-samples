import { input } from "pynput";

const items = new Map();
const price = 0.0;
const quantity = 0;

for (let i = 0; i < 3; i++) {
	const price = Number(await input(`Enter the price of item ${i + 1}: `));
	const quantity = Number(await input(`Enter the quantity of item ${i + 1}: `));

	items.set(i, { price, quantity });
}

let subtotal = 0.0;
const taxRate = 5.5;

for (const [_, value] of items) {
	subtotal += value.price * value.quantity;
}

const tax = (taxRate / 100) * subtotal;
console.log(`Subtotal: $${subtotal.toFixed(2)}`);
console.log(`Tax: $${tax.toFixed(2)}`);
console.log(`Total: $${(subtotal + tax).toFixed(2)}`);
