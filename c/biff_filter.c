#include <stdio.h>
#include <stdlib.h>

int main(void) {
  char message[100] = {};

  printf("Enter message: ");
  fgets(message, 100, stdin);

  // switch the characters over the loop.
  for (int i = 0; i < sizeof(message) / sizeof(message[0]); i++) {
    switch (message[i]) {
    case 'A':
      message[i] = '4';
      break;
    case 'B':
      message[i] = '8';
      break;
    case 'E':
      message[i] = '3';
      break;
    case 'I':
      message[i] = '1';
      break;
    case 'O':
      message[i] = '0';
      break;
    case 'S':
      message[i] = '5';
      break;
    default:
      message[i] = message[i];
    }
  }

  printf("In BIFF-speak: %s!!!!!!!!!!\n", message);

  return EXIT_SUCCESS;
}
