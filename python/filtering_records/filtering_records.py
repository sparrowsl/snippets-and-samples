import sqlite3
from typing import Any
# from typing import Any


with sqlite3.connect("./employees.db") as connection:
    cursor = connection.cursor()

search = input("Enter a search string: ")

rows = cursor.execute(
    "SELECT * FROM employees WHERE first_name LIKE ? or last_name LIKE ?",
    (f"%{search}%", f"%{search}%"),
).fetchall()


employees: list[dict[str, Any]] = []

for name, surname, position, date in rows:
    employees.append(
        {
            "name": name,
            "surname": surname,
            "position": position,
            "date": date,
        }
    )

print("Results:")
print(f"{'Name':20}| {'Position':12}| Separation Date")
print(f"{'-'*20}| {'-'*12}| {'-'*15}")
for emp in employees:
    name = f"{emp["name"]} {emp["surname"]}"
    print(f"{name:20}| {emp["surname"]:12}| {emp["date"] or ''}")
