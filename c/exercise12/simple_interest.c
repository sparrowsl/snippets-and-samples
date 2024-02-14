#include <stdio.h>
#include <stdlib.h>

int main(void) {
  float principal = 0.0;
  printf("Enter the principal: ");
  scanf("%f", &principal);

  float rate = 0.0;
  printf("Enter the rate of interest: ");
  scanf("%f", &rate);

  int years = 0;
  printf("Enter the years: ");
  scanf("%d", &years);

  float interest = (principal * (1 + (float)years * (rate / 100)));

  printf("After %d years at %.1f%%, the investment will be worth $%.2f.\n",
         years, rate, interest);

  return EXIT_SUCCESS;
}
