import fs from "node:fs";
import { input } from "pynput";

const input_file = "./input.txt";
const contents = fs.readFileSync(input_file, "utf8");

let output_file = "./output.txt";

const output = await input("Enter output file: ");
if (output.trim() !== "") {
	output_file = output;
}

const replaced = contents.replace(/utilize/g, "use");

fs.writeFileSync(output_file, replaced);
