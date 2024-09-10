import { input } from "pynput";

const value = await input("What is the input string? ");

if (value.trim().length === 0) {
	console.log("Please enter something.");
	process.exit();
}

const total = count_chars(value);
console.log(`"${value}" has ${total} characters.`);

/**
 * @param {string} sentence
 * @returns {number}
 */
function count_chars(sentence) {
	return sentence.length;
}
