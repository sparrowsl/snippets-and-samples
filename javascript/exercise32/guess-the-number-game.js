import { input } from "pynput";

const difficulty = Object.freeze({ 1: 10, 2: 100, 3: 1_000 });
let guessCount = 0;

console.log("Let's play Guess the Number.");
const level = await input("Pick a difficulty level (1, 2, 3): ");

const secret = Math.ceil(Math.random() * difficulty[level]);

let guess = await input("I have my number. What's your guess? ");

while (true) {
  guessCount += 1;

  if (guess > secret) {
    guess = await input("Too high. Guess again: ");
  } else if (guess < secret) {
    guess = await input("Too low. Guess again: ");
  } else {
    console.log(`You got it in ${guessCount} guesses!`);
    break;
  }
}
