def count_chars(sentence: str) -> int:
    return len(sentence)


value: str = input("What is the input string? ").strip()
total = count_chars(value)
print(f"{value} has {total} characters.")
