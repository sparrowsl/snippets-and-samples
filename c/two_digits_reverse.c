#include <stdio.h>
#include <stdlib.h>

int main(void) {
  int number = 0;

  printf("Enter a two-digit number: ");
  scanf("%d", &number);

  // get last number formula
  int last_number = number % 10;

  // get first number formula
  int first_number = number / 10;

  printf("The reversal is: %d%d\n", last_number, first_number);

  return EXIT_SUCCESS;
}
