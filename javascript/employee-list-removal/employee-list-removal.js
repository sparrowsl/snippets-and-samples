import fs from "node:fs";
import { input } from "pynput";

let employees = read_employees().split("\n");
display_employees(employees);

const name = await input("Enter an employee name to remove: ");
const nameIdx = employees.indexOf(name);

if (nameIdx < 0) {
	console.error(`"${name}" is not found`);
	process.exit(0);
}

employees = Array().concat(
	employees.slice(0, nameIdx),
	employees.slice(nameIdx + 1),
);

display_employees(employees);

function read_employees() {
	const filename = "./employees.txt";

	const buffer = fs.readFileSync(filename);
	return buffer.toString().trim();
}

/** @param {string[]} names */
function display_employees(names) {
	console.log(`There are ${names.length} employees:`);
	for (const name of names) {
		console.log(name);
	}
	console.log();
}
