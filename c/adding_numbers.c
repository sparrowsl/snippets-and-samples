#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>

int main(void) {
  int amount = 0;
  printf("Numbers to add? ");
  scanf("%d", &amount);

  int numbers[amount];
  int total = 0;

  for (int i = 0; i < amount; i++) {
    printf("Enter a number: ");
    scanf("%d", &numbers[i]);

    if (!isdigit(numbers[i])) {
      total += numbers[i];
    }
  }

  printf("The total is %d\n", total);

  return EXIT_SUCCESS;
}
