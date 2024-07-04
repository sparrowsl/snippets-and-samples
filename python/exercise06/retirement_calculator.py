from datetime import datetime

age = int(input("What is your current age? "))
retire_age = int(input("At what age would you like to retire? "))

# quit program if user is already retire
if age >= retire_age:
	print("You are already retired!!")
	exit(0)

current_year = datetime.now().year

print(f"You have {retire_age - age} years left until you can retire.")
print(f"It's {current_year}, so you can retire in {current_year + (retire_age - age)}.")
