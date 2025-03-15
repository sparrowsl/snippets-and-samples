import { input } from "pynput";

const euros = Number(await input("How many euros are you exchanging? "));
const exchange_rate = Number(await input("What is the exchange rate? "));

const to_dollars = (euros.toFixed(2) * exchange_rate) / 100;
console.log(
	`${euros} euros at an exchange rate of ${exchange_rate} is\n${to_dollars.toFixed(2)} U.S. dollars.`,
);
