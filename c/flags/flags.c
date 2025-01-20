#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "lexer.h"

typedef struct {
  char *key;
  char *value;
} Map;

void map_update(Map array[], char *string) {
  char *temp = "";
  int idx = 0; // current array in map to assign.

  for (int i = 0; i < (int)strlen(string); i++) {
    if (string[i] == ' ') {
      array[idx].key = temp;
      idx += 1;
      temp = ""; // reset the temp.
      continue;
    }

    temp += string[i];
    puts(temp);
  }
}

int str_split(Map array[], char *string) {
  int idx = 0;
  char *temp = "";

  for (int i = 0, len = (int)strlen(string); i < len; i++) {
    if (string[i] == ' ') {
      array[idx] = (Map){.key = temp};

      idx += 1;
      continue;
    }
  }

  // TODO: re-adjust implementation to not add 1 at the end for edge cases.
  return idx + 1;
}

int main(void) {
  /* char *args = "-name=jack -valid -dsn=file://flags.db"; */
  char *args = "-name=jack -valid=true";

  Lexer lex = new_lexer(args);

  printf("%c", lex.character);
  for (int i = 0, len = (int)strlen(args); i < len; i++) {
    printf("%c", lexer_peek(&lex));
    lexer_read_char(&lex);
  }

  /* printf("=> %s\n", lex.source); */
  /* printf("=> %d\n", lex.position); */

  return EXIT_SUCCESS;
}

/*void flag_parse(char *args) {}*/
