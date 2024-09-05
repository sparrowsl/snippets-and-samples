import { input } from "pynput";

const name = await input("What is your name? ")
console.log(`Hello, ${name}, nice to meet you!`)
