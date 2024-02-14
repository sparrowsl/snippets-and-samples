#include <math.h>
#include <stdio.h>
#include <stdlib.h>

struct item {
  float price;
  int quantity;
};

int main(void) {
  // make an array of items - length 3
  struct item items[3] = {};

  for (int i = 0; i < 3; i++) {
    float price = 0.0;
    int quantity = 0;

    printf("Enter the price of item %d: ", i + 1);
    scanf("%f", &price);

    printf("Enter the quantity of item %d: ", i + 1);
    scanf("%d", &quantity);

    // change value of the current item
    items[i].quantity = quantity;
    items[i].price = floorf(price);

    printf("\n"); // blank line after each user input
  }

  float subtotal = 0.0;
  const float tax_rate = 5.5;

  // sizeof(items) / sizeof(items[0]) - calculates the length of the items
  for (int i = 0; i < sizeof(items) / sizeof(items[0]); i++) {
    subtotal += (float)items[i].price * items[i].quantity;
  }

  float tax = (tax_rate / 100) * subtotal;
  printf("Subtotal: $%.2f\n", subtotal);
  printf("Tax: $%.2f\n", tax);
  printf("Total: $%.2f\n", subtotal + tax);

  return EXIT_SUCCESS;
}
