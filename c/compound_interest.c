#include <stdio.h>
#include <stdlib.h>

int main(void) {
  float principal = 0.0;
  printf("What is the principal amount? ");
  scanf("%f", &principal);

  float rate = 0.0;
  printf("What is the rate? ");
  scanf("%f", &rate);

  int years = 0;
  printf("What is the number of years? ");
  scanf("%d", &years);

  int compound_years = 0;
  printf("What is the number of times the interest is compounded per year? ");
  scanf("%d", &compound_years);

  float investment = principal * ((1 + (rate / (float)compound_years)) *
                                  (float)compound_years * years);

  printf("$%.2f invested at %.2f%% for %d years\ncompounded %d times per year "
         "is $%.2f.\n",
         principal, rate, years, compound_years, investment);

  return EXIT_SUCCESS;
}
