import { input } from "pynput";

const length = Number(await input("What is the length of room in feet? "));
const width = Number(await input("What is the width of room in feet? "));

const square_meters = Number(lengtNumber) * Number(width) * 0.09290304;

console.log(`You entered dimensions of ${length} feet by ${width} feet.`);
console.log(
	`The area is\n${length * width} square feet\n${square_meters.toFixed(3)} square meters`,
);
