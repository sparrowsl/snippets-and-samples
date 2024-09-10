import { input } from "pynput";

const password = await input("What is the password? ");

if (password === "abc$123") {
	console.log("Welcome!");
} else {
	console.log("I don't know you.");
}
