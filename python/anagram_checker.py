print("Enter two strings and I'll tell you if they are anagrams:")

first_string = input("Enter the first string: ")
second_string = input("Enter the second string: ")


def is_anagram(string1: str, string2: str) -> bool:
    # check if the strings have same length
    if len(string1) != len(string2):
        return False

    for i in range(len(string1)):
        if string1[i] not in string2:
            return False

    return True


if is_anagram(first_string, second_string):
    print(f"{first_string} and {second_string} are anagrams.")
else:
    print(f"{first_string} and {second_string} are not anagrams.")
