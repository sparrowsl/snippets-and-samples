typedef struct {
  char *source;
  int position;
  char character;
} Lexer;


Lexer new_lexer(char *);

char lexer_peek(Lexer *);

void lexer_read_char(Lexer *);

void lexer_skip_whitespaces(Lexer *);

