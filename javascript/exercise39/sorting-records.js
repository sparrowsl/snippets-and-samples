import Database from "better-sqlite3";

const db = new Database("records.db", { fileMustExist: true });
db.pragma("journal_mode = WAL");

/** @type [{first_name: string, last_name: string, position: string, date: string}] rows */
const rows = db.prepare("SELECT * FROM employees").all();
db.close(); // close db

rows.sort((a, b) => a.last_name.localeCompare(b.last_name));

console.log(
	`${"Name".padEnd(20)} | ${"Position".padEnd(20)} | Separation Date`,
);
console.log("-".repeat(60));
for (const row of rows) {
	const fullname = `${row.first_name} ${row.last_name}`;
	console.log(
		`${fullname.padEnd(20)} | ${row.position.padEnd(20)} | ${row.date || ""}`,
	);
}
