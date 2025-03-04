#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main(void) {
  srand(time(NULL));

  int difficulty_level = 0;
  int user_guess = 0;

  puts("Let's play Guess the Number.");
  printf("Pick a difficulty level (1, 2, or 3): ");
  scanf("%d", &difficulty_level);

  int max_difficulty = 10;
  switch (difficulty_level) {
  case 1:
    max_difficulty = 10;
    break;
  case 2:
    max_difficulty = 100;
    break;
  case 3:
    max_difficulty = 1000;
    break;
  }

  int secret = (rand() % max_difficulty) + 1;
  int guess_count = 1;

  printf("I have my number. What's your guess? ");
  while (true) {
    scanf("%d", &user_guess);

    if (user_guess < secret) {
      printf("Too low. Guess again: ");
    } else if (user_guess > secret) {
      printf("Too high. Guess again: ");
    } else {
      printf("You got it in %d guesses!!\n", guess_count);
      puts("Goodbye.");
      break;
    }

    guess_count += 1;
  }

  return EXIT_SUCCESS;
}
