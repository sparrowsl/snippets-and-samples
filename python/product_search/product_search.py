import json

products: list[dict[str, str | float]] = []


def search_product(name: str):
    for product in products:
        if str(product["name"]).lower() == name.lower():
            return product

    return None


def display_product(product: dict[str, str | float]):
    name, price, quantity = product.values()

    print(f"Name: {name}")
    print(f"Price: ${float(price):.2f}")
    print(f"Quantity on hand: {int(quantity)}")


def add_product(name: str):
    price = input("Product price: ").strip()
    quantity = input("Product quantity: ").strip()

    products.append({"name": name, "price": float(price), "quantity": int(quantity)})


with open("./products.json") as json_file:
    contents = json_file.read()
    products = json.loads(contents)["products"]


while True:
    product_name = input("What is the product name? ").strip()

    found = search_product(product_name)
    if not found:
        print("Sorry, that product was not found in out inventory.")
        response = input("Want to add the product? (y/n)? ")

        if response.strip() == "y":
            add_product(product_name)

        print()
        continue

    display_product(found)
    break
