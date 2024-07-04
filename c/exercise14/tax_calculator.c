#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
  float amount = 0.0;
  char state[3] = {};
  float tax = 5.5;

  printf("What is the order amount? ");
  scanf("%f", &amount);

  printf("What is the state? ");
  scanf("%s", state);

  // Check the amount and state condition
  if (strcmp(state, "WI")) {
    printf("The subtotal is $%.2f\n", amount);
    printf("The tax is $%.2f\n", (amount * tax) / 100);
    printf("The total is $%.2f\n", amount + ((amount * tax) / 100));
    return EXIT_FAILURE;
  }

  printf("The total is $%.2f\n", amount);
  return EXIT_SUCCESS;
}
