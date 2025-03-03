with open("./names.txt") as file:
    contents = file.read().strip()

sorted = sorted(contents.split("\n"))

print(f"Total of {len(sorted)} names")
print("-" * 18)

for name in sorted:
    print(name)
