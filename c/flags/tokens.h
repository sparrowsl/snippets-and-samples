#include "lexer.h"

typedef enum {
  ident,
  equal,
  eof,
} TokenType;

typedef struct {
  TokenType type;
  char *literal;
} Token;


/* Token new_token(TokenType, char); */

Token token_next_token(Lexer *);
