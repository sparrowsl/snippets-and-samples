principal = int(input("What is the principal amount? "))
rate = float(input("What is the rate? "))
years = int(input("What is the number of years? "))
compound_years = int(
    input("What is the number of times the interest is compounded per year? ")
)

investment = float(principal) * (
    (1 + ((rate / float(100)) / float(compound_years))) * float(compound_years * years)
)

print(
    f"${principal} invested at {rate:.2f} for {years} years\ncompounded {compound_years} times per year is ${investment:.2f}."
)
