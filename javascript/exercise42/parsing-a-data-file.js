import fs from "node:fs";
import Papa from "papaparse";

const csv_data = fs.readFileSync("./data.csv", "utf8").trim();
const parsed_data = Papa.parse(csv_data).data;

console.log(`${"Last".padEnd(10)} ${"First".padEnd(10)} Salary`);
console.log("-".repeat(28));
for (const row of parsed_data) {
	/** @type {string[]} */
	const [last, first, salary] = row;

	console.log(
		`${last.padEnd(10)} ${first.padEnd(10)} $${Intl.NumberFormat().format(salary)}`,
	);
}
