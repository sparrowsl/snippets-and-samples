#include <ctype.h>

#include "tokens.h"

Token new_token(TokenType type, char character) {
  return (Token){
      .type = type,
      .literal = &character,
  };
}

Token token_next_token(Lexer *lexer) {
  Token token;

  lexer_skip_whitespaces(lexer);

  switch (lexer->character) {
  case '=':
    token = new_token((TokenType)equal, lexer->character);
    break;
  case '\0':
    token = (Token){.type = eof, .literal = ""};
    break;
  default: {
    char tmp[100];
    int i = 0;
    // loop through each char checking if its a letter.
    while (isalpha(lexer->character) == 1) {
      tmp[i] = lexer->character;
      i++;
    }

    token = (Token){.type = ident, .literal = tmp};
  }
  }

  return token;
}
