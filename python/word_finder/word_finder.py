output_file = input("Name of output file: ")

with open("./input.txt") as input:
    contents = input.read()
    replaced = contents.replace("utilize", "use")

with open(f"./{output_file}", "w") as output:
    output.write(replaced)
