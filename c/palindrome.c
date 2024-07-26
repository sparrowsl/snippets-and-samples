// Palindrome for 5 numbers
#include <stdio.h>
#include <stdlib.h>

int main(void) {
  int digits = 0;

  printf("%s", "digits: ");
  scanf("%d", &digits);

  // printf("%d - %d\n\n", digits / 10000, digits % 10);
  // printf("%d - %d\n\n", digits / 1000 % 10, digits % 100 / 10);
  // printf("%d - %d\n\n", digits / 100 % 10, digits % 1000 / 100);

  if (digits / 10000 != digits % 10) {
    printf("%d is not a palindrome\n", digits);
    return EXIT_SUCCESS;
  }

  if (digits / 1000 % 10 != digits % 100 / 10) {
    printf("%d is not a palindrome\n", digits);
    return EXIT_SUCCESS;
  }

  if (digits / 100 % 10 != digits % 1000 / 100) {
    printf("%d is not a palindrome\n", digits);
    return EXIT_SUCCESS;
  }

  printf("%d is a palindrome\n", digits);
  return EXIT_SUCCESS;
}
