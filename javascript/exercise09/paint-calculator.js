import { input } from "pynput";

const length = Number(await input("Enter the length in feet? "));
const width = Number(await input("Enter the width in feet? "));

console.log(
	`You will need to purchase ${(length * width) / 350} gallons of\npaint to cover ${length * width} square feet.`,
);
