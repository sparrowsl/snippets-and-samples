#include <ctype.h>
#include <stdbool.h>

#include "tokens.h"

Token new_token(TokenType type, char character) {
  return (Token){
      .type = type,
      .literal = &character,
  };
}

Token token_next_token(Lexer *lexer) {
  lexer_skip_whitespaces(lexer);

  Token token;
  switch (lexer->character) {
  case '=':
    token = new_token((TokenType)equal, lexer->character);
    lexer_advance(lexer);
    break;
  case '\0':
    token = (Token){.type = eof, .literal = ""};
    break;
  default: {
    char buffer[256];
    int idx = 0;

    // loop through each char checking if its a letter, while taking into
    // account the minus sign.
    while (lexer->character == '-' || isalpha(lexer->character)) {
      buffer[idx] = lexer->character;
      idx += 1;
      lexer_advance(lexer);
    }

    buffer[idx] = '\0';
    token = (Token){.type = ident, .literal = buffer};
  }
  }

  return token;
}
