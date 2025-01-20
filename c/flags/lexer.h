typedef struct {
  char *source;
  int position;
  char character;
} Lexer;


Lexer new_lexer(char *source);

char lexer_peek(Lexer *lexer);

void lexer_read_char(Lexer *lexer);

void lexer_skip_whitespaces(Lexer *lexer);

