#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
  char sentence[50];
  printf("What is the input string? ");
  fgets(sentence, sizeof(sentence), stdin);

  if (strlen(sentence) == 0) {
    printf("Please enter something.\n");
    return EXIT_SUCCESS;
  }

  printf("\"%s\" has %lu characters.\n", sentence, strlen(sentence));

  return EXIT_SUCCESS;
}
