age = input("What is your age? ")

try:
    age = int(age)
except ValueError:
    print("enter a valid number")
    exit(0)

if age < 1 or age < 16:
    print("You are not old enough to legally drive")
else:
    print("You are old enough to legally drive")
