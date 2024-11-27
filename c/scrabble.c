#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int get_face_value(char letter) {
  switch (toupper(letter)) {
  case 'A':
  case 'E':
  case 'I':
  case 'L':
  case 'N':
  case 'O':
  case 'R':
  case 'S':
  case 'T':
  case 'U':
    return 1;
  case 'D':
  case 'G':
    return 2;
  case 'B':
  case 'C':
  case 'M':
  case 'P':
    return 3;
  case 'F':
  case 'H':
  case 'V':
  case 'W':
  case 'Y':
    return 4;
  case 'K':
    return 5;
  case 'J':
  case 'X':
    return 8;
  case 'Q':
  case 'Z':
    return 10;
  default:
    return 0;
  }
}

int main(void) {
  char word[20];

  printf("Enter a word: ");
  scanf("%s", word);

  int total_value = 0;
  for (int i = 0, len = (int)strlen(word); i < len; i++) {
    total_value += get_face_value(word[i]);
  }

  printf("Scrabble value: %d\n", total_value);

  return EXIT_SUCCESS;
}
