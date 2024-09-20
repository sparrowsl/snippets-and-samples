import fs from "node:fs";

const contents = fs.readFileSync("./names.txt", "utf8").trim().split("\n");

contents.sort((a, b) => a.localeCompare(b));

console.log(`Total of ${contents.length} names`);
console.log("-".repeat(18));
for (const name of contents) {
	console.log(name);
}
