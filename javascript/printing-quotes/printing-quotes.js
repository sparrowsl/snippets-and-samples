import { input } from "pynput";

const quote = await input("What is the quote? ");
const author = await input("Who said it? ");

const result = print_quote(author, quote);
console.log(result);

/**
 * @param {string} quote
 * @param {string} author
 * @returns {string}
 */
function print_quote(author, quote) {
	return `${author} says, \"${quote}\"`;
}
