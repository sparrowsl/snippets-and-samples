import { input } from "pynput";

const age = await input("What is your age? ");

if (age < 16) {
	console.log("You are not old enough to legally drive");
} else {
	console.log("You are old enough to legally drive");
}
