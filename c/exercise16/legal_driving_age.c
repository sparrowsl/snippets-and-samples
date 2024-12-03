#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main(void) {
  int age = 0;

  printf("What is your age? ");
  scanf("%d", &age);

  if (age < 16) {
    puts("You are not old enough to legally drive.");
  } else {
    puts("You are old enough to legally drive.");
  }

  return EXIT_SUCCESS;
}
