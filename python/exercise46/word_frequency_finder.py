from flask import Flask

app = Flask(__name__)

store: dict[str, int] = dict()


def count_words(words: list[str]):
    for word in words:
        store[word] = words.count(word)


with open("./words.txt") as file:
    contents = file.read().replace("\n", " ").strip()

    stripped = contents.translate(
        {ord(i): None for i in ",;:_[]?!@\"'.—”“’)(&*%$#-|I\n\t"}
    )
    # stripped = contents.strip(",;:_[]?!@\"'.—”“’)(&*%$#-|I\r\n")
    splitted = stripped.split(" ")


@app.get("/")
def index():
    count_words(splitted)
    return "<h1>Hello again Flask!!</h1>"
