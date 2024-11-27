#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int calculate_average(char *);

int main(void) {
  char sentence[100];

  printf("Enter a sentence: ");
  fgets(sentence, 100, stdin);
  sentence[strcspn(sentence, "\n\r")] = 0;

  int avg = calculate_average(sentence);
  printf("Average word length: %d", avg);

  return EXIT_SUCCESS;
}

int calculate_average(char *desc) {
  // buggy version here
  int letters = 0;
  int words = 0;

  for (int i = 0, len = strlen(desc); i < len; i++) {
    if (desc[i] == ' ') {
      words += 1;
    } else if (desc[i] != ' ') {
      letters += 1;
    }
  }

  return letters / words;
}
