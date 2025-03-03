#include <stdio.h>
#include <stdlib.h>

int main(void) {
  double height = 0.0;
  double weight = 0.0;
  double bmi = 0;

  printf("Enter height: ");
  scanf("%lf", &height);

  printf("Enter weight: ");
  scanf("%lf", &weight);

  bmi = (weight / (height * height)) * 703;

  printf("Your BMI is %.2f\n", bmi);
  if (bmi >= 18.5 && bmi <= 25) {
    puts("You are within the ideal weight range.");
  } else {
    puts("You are overweight. You should see your doctor.");
  }

  return EXIT_SUCCESS;
}
