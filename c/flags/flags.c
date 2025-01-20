#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* #include "lexer.h" */
#include "tokens.h"

int main(void) {
  /* char *args = "-name=jack -valid -dsn=file://flags.db"; */
  char *args = "-name=jack -valid=true";

  /* Token tokens[10] = {}; */

  Lexer lex = new_lexer(args);

  for (int i = 0, len = (int)strlen(args); i < len; i++) {
    lexer_read_char(&lex);
    printf("%c", lex.character);
    lexer_skip_whitespaces(&lex);
  }

  /* printf("=> %s\n", lex.source); */
  /* printf("=> %d\n", lex.position); */

  return EXIT_SUCCESS;
}
