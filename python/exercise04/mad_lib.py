def create_mad_libs(noun: str, verb: str, adjective: str, adverb: str) -> str:
    return f"Do you {verb} your {adjective} {noun} {adverb}? That's hilarious"


noun = input("Enter a noun: ")
verb = input("Enter a verb: ")
adjective = input("Enter an adjective: ")
adverb = input("Enter an adverb: ")

result = create_mad_libs(noun, verb, adjective, adverb)
print(result)
