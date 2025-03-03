import fs from "node:fs";
import { input } from "pynput";

const contents = fs.readFileSync("./products.json", "utf8");

/** @type {{name: string, price: number, quantity: number}[]} */
const products = JSON.parse(contents).products;

while (true) {
	const product_name = await input("What is the product name? ");
	const found = find_product(product_name);

	if (found) {
		display_product(found);
		break;
	}
	console.log("Sorry, that product was not found in our inventory.");

	const valid = await input("Add the product? (y/n)? ");
	if (valid.trim() === "y") {
		await add_product(product_name);
	}
}

/** @param {{name: string, price: number, quantity: number}} product  */
function display_product(product) {
	console.log(`Name: ${product.name}`);
	console.log(`Price: $${product.price.toFixed(2)}`);
	console.log(`Quantity on hand: ${product.quantity}`);
}

/** @param {string} name  */
function find_product(name) {
	return products.find(
		(product) => product.name.toLowerCase() === name.toLowerCase(),
	);
}

/** @param {string} name */
async function add_product(name) {
	const price = await input("Price: ", { convert: "float" });
	const quantity = await input("Quantity: ", { convert: "int" });
	products.push({ name, price, quantity });
	console.log();
}
