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

const monthIndex = await input("Please enter the number of the month: ");
console.log(`The name of the month is ${months[monthIndex - 1]}`);
