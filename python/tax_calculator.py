tax = 5.5

amount = float(input("What is the order amount? "))
state = input("What is the state? ")

# Check the amount and state condition
if state.upper() == "WI":
    print(f"The subtotal is ${amount:.2f}")
    print(f"The tax is ${((amount * tax) / 100):.2f}")
    print(f"The total is ${(amount + ((amount * tax) / 100)):.2f}")
    exit(0)


print(f"The total is ${amount:.2f}")
