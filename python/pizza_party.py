number_of_people = input("How many people? ")
try:
    people = int(number_of_people)
except ValueError:
    print("Please enter a valid number!!")
    exit(0)

number_of_pizzas = input("How many pizzas do you have? ")
try:
    pizzas = int(number_of_pizzas)
except ValueError:
    print("Please enter a valid number!!")
    exit(0)

leftovers = people % pizzas
print(f"{people} people with {pizzas} pizzas")
print(f"Each person gets {(people / pizzas) / 2} pieces of pizza.")
print(f"There are {leftovers} leftovers pieces.")
