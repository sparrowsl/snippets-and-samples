import { input } from "pynput"

const principal = Number(await input("Enter the principal: "))
const rate = Number(await input("Enter the rate of interest: "))
const years = Number(await input("Enter the years: "))

const interest = (principal.toFixed(2) * (1 + (years * (rate / 100))))

console.log(`After ${years} years at ${rate}%, the investment will be worth $${interest}.`)
