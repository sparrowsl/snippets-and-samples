// Palindrome for words
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
  char *word = malloc(50);

  printf("%s", "word: ");
  fgets(word, sizeof(word), stdin);

  for (int i = 0, len = strlen(word) - 1; i < len; i++) {
    char last = word[len - (i + 1)];

    if (word[i] != last) {
      word[strcspn(word, "\n")] = 0; // remove the \n from word
      printf("%s is not a palindrome!!\n", word);
      return EXIT_FAILURE;
    }
  }

  word[strcspn(word, "\n")] = 0;
  printf("%s is a palindrome\n", word);

  return EXIT_SUCCESS;
}
