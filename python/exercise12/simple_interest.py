principal = int(input("Enter the principal: "))
rate = float(input("Enter the rate of interest: "))
years = int(input("Enter the years: "))

interest = float(principal) * (1 + float(years) * (rate / 100))

print(f"After {years} years at {rate:.1f}%, the investment will be worth ${interest}.")
