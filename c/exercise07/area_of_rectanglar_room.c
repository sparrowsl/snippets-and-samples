#include <stdio.h>
#include <stdlib.h>

int main(void) {
  float length = 0;
  float width = 0;

  printf("What is the length of room in feet? ");
  scanf("%f", &length);

  printf("What is the width of room in feet? ");
  scanf("%f", &width);

  float square_meters = (length * width) * 0.09290304;
  printf("You entered dimensions of %.2f feet by %.2f feet.\n", length, width);
  printf("The area is\n%.2f square feet\n%.3f square meters\n", length * width,
         square_meters);

  return EXIT_SUCCESS;
}
