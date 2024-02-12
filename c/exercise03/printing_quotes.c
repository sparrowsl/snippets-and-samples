#include <stdio.h>
#include <stdlib.h>

int main(void) {
  char quote[50] = {};
  char author[30] = {};

  printf("What is the quote? ");
  fgets(quote, sizeof(quote), stdin);

  printf("Who said it? ");
  fgets(author, sizeof(author), stdin);

  printf("%s says, \"%s\"\n", author, quote);

  return EXIT_SUCCESS;
}
