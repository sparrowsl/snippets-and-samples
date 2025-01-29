#include <string.h>

typedef struct {
  char *source;
  int position;
  char character;
} Lexer;

Lexer new_lexer(char *);
char lexer_peek(Lexer *);
void lexer_advance(Lexer *);
void lexer_read_char(Lexer *);
void lexer_skip_whitespaces(Lexer *);

Lexer new_lexer(char *source) {
  Lexer lexer = {
      .source = source,
      .position = 0,
      .character = source[0],
  };

  return lexer;
}

char lexer_peek(Lexer *lexer) { return lexer->source[lexer->position + 1]; }

void lexer_advance(Lexer *lexer) {
  lexer->position += 1;
  lexer->character = lexer->source[lexer->position];
}

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
    lexer_advance(lexer);
  }
}
