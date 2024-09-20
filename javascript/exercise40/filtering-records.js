import Database from "better-sqlite3";
import { input } from "pynput";

const db = new Database("./records.db", { fileMustExist: true });

const search = await input("Enter a search string: ");

/** @type [{
 * first_name: string,
 * last_name: string,
 * position: string,
 * date: string
 * }] rows */
const rows = db
	.prepare("SELECT * FROM employees WHERE first_name like ?")
	.all(`${search}%`);
db.close();

// console.log(rows)

console.log(
	`${"Name".padEnd(20)} | ${"Position".padEnd(20)} | Separation Date`,
);
console.log("-".repeat(60));
for (const row of rows) {
	const fullname = `${row.first_name} ${row.last_name}`.padEnd(20);
	const position = row.position.padEnd(20);
	const date = row.date || "";

	console.log(`${fullname} | ${position} | ${date}`);
}
