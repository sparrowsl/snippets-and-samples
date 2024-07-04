def print_quote(author:str, quote: str)-> str:
  return f"{author} says, \"{quote}\""

quote: str = input("What is the quote? ")
author: str = input("Who said it? ")

result = print_quote(author, quote)
print(result)
