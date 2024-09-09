import { input } from "pynput";

const names = [];

while (true) {
  const name = await input("Enter a name: ");

  if (name === "") {
    break;
  }

  names.push(name);
}

const winner = pick_random_winner(names);
console.log(`The winner is ${winner}`);

/** @param {string} names */
function pick_random_winner(names) {
  const random_number = Math.ceil(Math.random() * names.length);
  return names[random_number];
}
