#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int count_vowels(char *);

int main(void) {
  char sentence[100];
  printf("Enter a sentence: ");
  fgets(sentence, 100, stdin);
  sentence[strcspn(sentence, "\n\r")] = 0;

  printf("Your sentence contains %d vowels.\n", count_vowels(sentence));

  return EXIT_SUCCESS;
}

int count_vowels(char *desc) {
  int amount = 0;

  for (int i = 0, len = strlen(desc); i < len; i++) {
    switch (tolower(desc[i])) {
    case 'a':
    case 'e':
    case 'i':
    case 'o':
    case 'u':
      amount += 1;
    }
  }

  return amount;
}
