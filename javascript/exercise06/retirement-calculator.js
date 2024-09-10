import { input } from "pynput";

const age = Number(await input("What is your current age? "));
const retireAge = Number(await input("At what age would you like to retire? "));

// quit program if user is already retire
if (age >= retireAge) {
	console.log("You are already retired!!");
	process.exit();
}

const currentYear = new Date().getFullYear();

console.log(`You have ${retireAge - age} years left until you can retire.`);
console.log(
	`It's ${currentYear}, so you can retire in ${currentYear + (retireAge - age)}.`,
);
