import { input } from "pynput"

const euros = Number(await input("How many euros are you exchanging? "))
const exchangeRate = Number(await input("What is the exchange rate? "))

const toDollars = (euros.toFixed(2) * exchangeRate) / 100
console.log(`${euros} euros at an exchange rate of ${exchangeRate} is\n${toDollars.toFixed(2)} U.S. dollars.`)
