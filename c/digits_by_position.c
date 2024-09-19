#include <stdio.h>
#include <stdlib.h>

int main(void) {
  int digits = 0;

  printf("%s", "Enter digits: ");
  scanf("%d", &digits);

  // first digit
  printf("%d ", digits / 10000);
  // second digit
  printf("%d ", (digits % 10000) / 1000);
  // third digit
  printf("%d ", (digits % 1000) / 100);
  // fourth digit
  printf("%d ", (digits % 100) / 10);
  // fifth digit
  printf("%d \n", digits % 10);

  return EXIT_SUCCESS;
}
