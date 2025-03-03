import csv


def format_row(row: list[str]):
    last_name, first_name, salary = row

    print(f"{last_name:10} {first_name:10} ${int(salary):0,}")


with open("./data.csv") as file:
    rows = csv.reader(file, delimiter=",")

    print(f"{"Last":10} {"First":10} Salary")
    print("-" * 30)
    for row in rows:
        format_row(row)
