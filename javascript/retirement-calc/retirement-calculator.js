import { input } from "pynput";

const age = Number(await input("What is your current age? "));
const retire_age = Number(
	await input("At what age would you like to retire? "),
);

// quit program if user is already retire
if (age >= retire_age) {
	console.log("You are already retired!!");
	process.exit();
}

const current_year = new Date().getFullYear();

console.log(`You have ${retire_age - age} years left until you can retire.`);
console.log(
	`It's ${current_year}, so you can retire in ${current_year + (retire_age - age)}.`,
);
