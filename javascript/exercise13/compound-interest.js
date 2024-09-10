import { input } from "pynput";

const principal = Number(await input("What is the principal amount? "));
const rate = Number(await input("What is the rate? "));
const years = Number(await input("What is the number of years? "));
const compoundYears = Number(
	await input(
		"What is the number of times the interest is compounded per year? ",
	),
);

const investment =
	principal * ((1 + rate / 100 / compoundYears) * compoundYears * years);

console.log(
	`$${principal} invested at ${rate.toFixed(2)} for ${years} years\ncompounded ${compoundYears} times per year is $${investment.toFixed(2)}.`,
);
