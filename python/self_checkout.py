items: list[dict[str, float | int]] = list()
price = 0.0
quantity = 0

for i in range(1, 4):
    price = float(input(f"Enter the price of item {i}: "))
    quantity = int(input(f"Enter the quantity of item {i}: "))

    items.append({"price": price, "quantity": quantity})


sub_total = 0.0
tax_rate = 5.5


for item in items:
    sub_total += item["price"] * item["quantity"]

tax = (tax_rate / 100) * sub_total
print(f"Subtotal: ${sub_total:.2f}")
print(f"Tax: ${tax:.2f}")
print(f"Total: ${(sub_total + tax):.2f}")
