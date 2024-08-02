euros = int(input("How many euros are you exchanging? "))
exchange_rate = float(input("What is the exchange rate? "))

to_dollars = (float(euros) * exchange_rate) / 100

print(
    f"{euros} euros at an exchange rate of {exchange_rate} is\n{to_dollars:.2f} U.S. dollars."
)
