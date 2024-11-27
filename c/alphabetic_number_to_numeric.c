#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
  char number[15];

  printf("Enter phone number: ");
  scanf("%s", number);

  int i = 0;
  while (i < (int)strlen(number)) {
    if (isdigit(number[i])) {
      printf("%c", number[i]);
    } else if (isalpha(number[i])) {
      switch (toupper(number[i])) {
      case 'A' ... 'C':
        printf("%d", 2);
        break;
      case 'D' ... 'F':
        printf("%d", 3);
        break;
      case 'G' ... 'I':
        printf("%d", 4);
        break;
      case 'J' ... 'L':
        printf("%d", 5);
        break;
      case 'M' ... 'O':
        printf("%d", 6);
        break;
      case 'P' ... 'S':
        printf("%d", 7);
        break;
      case 'T' ... 'V':
        printf("%d", 8);
        break;
      case 'W' ... 'Y':
        printf("%d", 9);
        break;
      default:
        printf("%d", number[i]);
      }
    } else {
      printf("%c", number[i]);
    }

    i += 1;
  }

  return EXIT_SUCCESS;
}
