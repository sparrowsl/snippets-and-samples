#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#include "tokens.c"

int main(void) {
  /* char *args = "-name=jack -valid -dsn=file://flags.db"; */
  char *args = "-name=jack -valid=true";

  Lexer lex = new_lexer(args);

  while (lex.character != '\0') {
    Token tk = token_next_token(&lex);
    lexer_advance(&lex);
    printf("token: %s\n", tk.literal);
  }

  return EXIT_SUCCESS;
}
