from pathlib import Path

filepath = Path("./words.txt")
store: dict[str, int] = dict()


def count_words(words: list[str]):
    for word in words:
        store[word] = words.count(word)


with open(filepath) as file:
    contents = file.read().replace("\n", " ").strip()

    stripped = contents.translate(
        {ord(i): None for i in ",;:_[]?!@\"'.—”“’)(&*%$#-|I\n\t"}
    )
    # stripped = contents.strip(",;:_[]?!@\"'.—”“’)(&*%$#-|I\r\n")
    splitted = stripped.split(" ")

count_words(splitted)

for key, value in store.items():
    print(f"{key}: {"*" * value}")
