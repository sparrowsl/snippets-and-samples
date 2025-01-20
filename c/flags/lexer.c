#include <string.h>

#include "lexer.h"

Lexer new_lexer(char *source) {
  Lexer lexer = {
      .source = source,
      .position = 0,
      .character = source[0],
  };

  return lexer;
}

char lexer_peek(Lexer *lexer) { return lexer->source[lexer->position + 1]; }

void lexer_read_char(Lexer *lexer) {
  if (lexer->position > (int)strlen(lexer->source)) {
    lexer->character = lexer->source[0];
  } else {
    lexer->character = lexer->source[lexer->position];
  }

  lexer->position += 1;
}

void lexer_skip_whitespaces(Lexer *lexer) {
  while (lexer->character == ' ' || lexer->character == '\n' ||
         lexer->character == '\t' || lexer->character == '\r') {
    lexer_read_char(lexer);
  }
}
