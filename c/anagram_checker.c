#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

bool is_anagram(char *, char *);

int main(void) {
  char word1[30];
  char word2[30];

  puts("Enter 2 strings to determine if they are anagrams.");
  printf("Enter first string: ");
  scanf("%s", word1);

  printf("Enter second string: ");
  scanf("%s", word2);

  if (is_anagram(word1, word2)) {
    printf("\"%s\" and \"%s\" are anagrams.\n", word1, word2);
  } else {
    printf("\"%s\" and \"%s\" are not anagrams.\n", word1, word2);
  }

  return EXIT_SUCCESS;
}

bool is_anagram(char *word1, char *word2) {
  if (strlen(word1) != strlen(word2)) {
    return false;
  }

  for (int i = 0; i < (int)strlen(word2); i++) {
    if (word1[0] == word2[i]) {
      return true;
    }
  }

  return false;
}
