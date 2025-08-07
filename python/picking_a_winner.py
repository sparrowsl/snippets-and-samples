import random

names: list[str] = []


def pick_winner(names: list[str]) -> str:
    num = random.randint(0, len(names) - 1)
    return names[num]


while True:
    name = input("Enter a name: ").strip()

    if name == "":
        break

    names.append(name)


winner = pick_winner(names)
print(f"The winner is... {winner}")
