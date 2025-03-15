import { input } from "pynput";

const noun = await input("Enter a noun: ");
const verb = await input("Enter a verb: ");
const adjective = await input("Enter an adjective: ");
const adverb = await input("Enter an adverb: ");

const result = create_mad_libs(noun, verb, adjective, adverb);
console.log(result);

/**
 * @param {string} noun
 * @param {string} verb
 * @param {string} adverb
 * @param {string} string
 * @param {string} adjective
 * @returns {string}
 */
function create_mad_libs(noun, verb, adjective, adverb) {
	return `Do you ${verb} your ${adjective} ${noun} ${adverb}? That's hilarious`;
}
