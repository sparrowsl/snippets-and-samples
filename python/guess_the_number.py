import random

levels = {1: 10, 2: 100, 3: 1000}
print("Let's play Guess the Number.")
difficulty = input("Pick a difficulty level (1, 2, or 3): ")


try:
    difficulty = int(difficulty)
    if difficulty not in [1, 2, 3]:
        print(f"{difficulty} is not a valid difficulty level!")
        exit(0)
except ValueError:
    print("Invalid difficulty level")
    exit(0)


secret = random.randrange(1, levels[difficulty])
guess_count = 0

while True:
    try:
        print(f"Guess a number between 1 and {levels[difficulty]}!")
        guess = int(input("I have my number. What's your guess? "))
    except ValueError:
        print("Invalid number")
        continue
    finally:
        guess_count += 1

    if guess > secret:
        print("Too high, Guess again")
    elif guess < secret:
        print("Too low, Guess again")
    else:
        print(f"You got it in {guess_count} guesses!")
        break

    print("")
