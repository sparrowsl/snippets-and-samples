// Guessing game
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int start_game(void);

int main(void) {
  srand(time(NULL));

  start_game();

  return EXIT_SUCCESS;
}

int start_game(void) {
  int guess = 0;
  int secret_number = 1 + (rand() % 1000);

  puts("I have a number between 1 and 1,000.");
  puts("can you guess my number?");
  printf("Please type your first guess.\n> ");

  while (guess != secret_number) {
    scanf("%d", &guess);

    if (guess > secret_number) {
      puts("Too high. Try again.");
    } else if (guess < secret_number) {
      puts("Too low. Try again.");
    } else {
      puts("Excellent!!! you guessed the number.");
    }
  }

  return EXIT_SUCCESS;
}
