#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
  char smallest[50];
  char largest[50];
  char word[50];

  while (true) {
    printf("Enter word: ");
    scanf("%s", word);

    if (strlen(word) == 4) {
      break;
    }

    if (strlen(smallest) == 0 || strlen(smallest) > strlen(word) ||
        strlen(smallest) == strlen(word)) {
      strcpy(smallest, word);
    } else if (strlen(largest) < strlen(word)) {
      strcpy(largest, word);
    }
  }

  printf("Smallest word: %s\n", smallest);
  printf("Largest word: %s\n", largest);

  return EXIT_SUCCESS;
}
