import { input } from "pynput";

console.log("Enter two strings and I'll tell you if they are anagrams:");
const first = await input("Enter the first string: ");
const second = await input("Enter the second string: ");

if (is_anagram(first, second)) {
	console.log(`"${first}" and "${second}" are anagrams.`);
}

/**
 * @param {string} first
 * @param {string} second
 */
function is_anagram(first, second) {
	if (first.length !== second.length) {
		return false;
	}

	for (const letter of first) {
		if (!second.includes(letter)) {
			return false;
		}
	}

	return true;
}
