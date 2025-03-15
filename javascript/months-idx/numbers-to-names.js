import { input } from "pynput";

const months = [
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
];

const month_index = await input("Please enter the number of the month: ");
console.log(`The name of the month is ${months[month_index - 1]}`);
