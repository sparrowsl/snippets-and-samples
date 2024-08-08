employees = [
    "John Smith",
    "Jackie Jackson",
    "Chris Jones",
    "Amanda Cullen",
    "Jeremy Goodwin",
]

print(f"There are {len(employees)} employees:")
for emp in employees:
    print(f"  - {emp}")

name = input("\nEnter an employee name to remove: ")

if name not in employees:
    print(f"{name} is not an employee...")
    exit(0)
employees.remove(name)

print(f"\nThere are {len(employees)} employees:")
for emp in employees:
    print(f"  - {emp}")
