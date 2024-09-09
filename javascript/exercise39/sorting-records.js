import Database from "better-sqlite3";
import { input } from "pynput";

const db = new Database("records.db", { fileMustExist: true });
db.pragma("journal_mode = WAL");

/** @type [{first_name: string, last_name: string, position: string, date: string}] rows */
const rows = db.prepare("SELECT * FROM employees").all();
rows.sort((a, b) => a.last_name.localeCompare(b.last_name));

console.log(
  `${"Name".padEnd(20)} | ${"Position".padEnd(20)} | Separation Date`,
);
console.log("-".repeat(60));
for (const user of rows) {
  console.log(
    `${`${user.first_name} ${user.last_name}`.padEnd(20)} | ${user.position.padEnd(20)} | ${user.date || ""}`,
  );
}

db.close();
