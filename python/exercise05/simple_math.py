def add(a: float, b: float) -> float:
    return a + b


def subtract(a: float, b: float) -> float:
    return a - b


def multiply(a: float, b: float) -> float:
    return a * b


def divide(a: float, b: float) -> float:
    if b == 0:
        return 0

    return a / b


first_number: int = int(input("What is the first number? "))
second_number: int = int(input("What is the second number? "))

sum = add(first_number, second_number)
difference = subtract(first_number, second_number)
product = multiply(first_number, second_number)
quotient = divide(first_number, second_number)

print("%v + %v = %v\n", first_number, second_number, sum)
print("%v - %v = %v\n", first_number, second_number, difference)
print("%v * %v = %v\n", first_number, second_number, product)
print("%v / %v = %v\n", first_number, second_number, quotient)
