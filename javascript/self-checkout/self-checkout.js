import { input } from "pynput";

const items = new Map();

for (let i = 0; i < 3; i++) {
	const price = Number(await input(`Enter the price of item ${i + 1}: `));
	const quantity = Number(await input(`Enter the quantity of item ${i + 1}: `));

	items.set(i, { price, quantity });
}

let subtotal = 0.0;
const tax_rate = 5.5;

for (const [_, value] of items) {
	subtotal += value.price * value.quantity;
}

const tax = (tax_rate / 100) * subtotal;
console.log(`Subtotal: $${subtotal.toFixed(2)}`);
console.log(`Tax: $${tax.toFixed(2)}`);
console.log(`Total: $${(subtotal + tax).toFixed(2)}`);
