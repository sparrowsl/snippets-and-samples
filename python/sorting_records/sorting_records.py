import sqlite3
from typing import Any


# Create the db
with sqlite3.connect("./employees.db") as connection:
    cursor = connection.cursor()

employees: list[dict[str, Any]] = []

# execute query
rows = cursor.execute("SELECT * FROM employees").fetchall()


for name, surname, position, date in rows:
    employees.append(
        {
            "name": name,
            "surname": surname,
            "position": position,
            "date": date,
        }
    )

# sorted(employees, key=lambda employee: employee["surname"])
employees.sort(key=lambda emp: emp["surname"])

print(f"{'Name':20}| {'Position':12}| Separation Date")
print(f"{'-'*20}| {'-'*12}| {'-'*15}")
for emp in employees:
    name = f"{emp["name"]} {emp["surname"]}"
    print(f"{name:20}| {emp["surname"]:12}| {emp["date"] or ''}")
