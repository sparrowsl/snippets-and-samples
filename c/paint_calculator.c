#include <stdio.h>
#include <stdlib.h>

int main(void) {
  int length = 0;
  printf("Enter the length in feet? ");
  scanf("%d", &length);

  int width = 0;
  printf("Enter the width in feet? ");
  scanf("%d", &width);

  printf("You will need to purchase %d gallons of\npaint to cover %d square "
         "feet.\n",
         (length * width) / 350, length * width);
  return EXIT_SUCCESS;
}
