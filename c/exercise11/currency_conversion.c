#include <stdio.h>
#include <stdlib.h>

int main(void) {
  float euros = 0.0;
  printf("How many euros are you exchanging? ");
  scanf("%f", &euros);

  float exchange_rate = 0.0;
  printf("What is the exchange rate? ");
  scanf("%f", &exchange_rate);

  float to_dollars = (euros * exchange_rate) / 100;
  printf("%.2f euros at an exchange rate of %.2f is\n%.2f U.S. dollars.\n",
         euros, exchange_rate, to_dollars);

  return EXIT_SUCCESS;
}
