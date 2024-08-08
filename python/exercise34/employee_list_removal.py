with open("./employees.txt") as file:
    employees = [file.strip("\n") for file in file.readlines()]


print(f"There are {len(employees)} employees:")
for emp in employees:
    print(f"  - {emp.strip()}")


name = input("\nEnter an employee name to remove: ")
if name not in employees:
    print(f"{name} is not an employee...")
    exit(0)

employees.remove(name)

print(f"\nThere are {len(employees)} employees:")
for emp in employees:
    print(f"  - {emp}")

with open("./employees.txt", "w") as file:
    file.writelines(employees)
