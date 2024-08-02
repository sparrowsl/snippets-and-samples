items: dict = {}
price = 0.0
quantity = 0

for i in range(3):
    price = float(input(f"Enter the price of item {i + 1}: "))
    quantity = int(input(f"Enter the quantity of item {i + 1}: "))

    items[i] = {"price": price, "quantity": float(quantity)}

print(items)

sub_total = 0.0
tax_rate = 5.5

for key in items.keys():
    sub_total += items[key]["price"] * items[key]["quantity"]

tax = (tax_rate / 100) * sub_total
print(f"Subtotal: ${sub_total:.2f}")
print(f"Tax: ${tax:.2f}")
print(f"Total: ${(sub_total + tax):.2f}")
