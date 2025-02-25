#include <stdbool.h>
#include <stdio.h>
#include <string.h>

#define MAX_FLAGS 20      // Max number of flags
#define MAX_KEY_LEN 50    // Max length of key
#define MAX_VALUE_LEN 100 // Max length of value

typedef struct {
  char key[MAX_KEY_LEN];
  char value[MAX_VALUE_LEN];
  bool is_boolean; // True if flag is a boolean switch
} Flag;

typedef struct {
  Flag flags[MAX_FLAGS];
  int count;
} Parser;

void parse_flags(Parser *parser, int argc, char *argv[]) {
  parser->count = 0;

  for (int i = 1; i < argc; i++) { // Start from 1, skip program name
    if (argv[i][0] == '-') {       // Ensure it's a flag
      if (parser->count >= MAX_FLAGS) {
        fprintf(stderr, "Error: Too many flags!\n");
        return;
      }

      char *equal_pos = strchr(argv[i], '=');
      if (equal_pos) {                            // It's a key=value flag
        size_t key_len = equal_pos - argv[i] - 1; // Exclude '-'
        size_t value_len = strlen(equal_pos + 1);

        if (key_len >= MAX_KEY_LEN || value_len >= MAX_VALUE_LEN) {
          fprintf(stderr, "Error: Key/Value too long!\n");
          continue;
        }

        strncpy(parser->flags[parser->count].key, argv[i] + 1, key_len);
        parser->flags[parser->count].key[key_len] = '\0';
        strncpy(parser->flags[parser->count].value, equal_pos + 1, value_len);
        parser->flags[parser->count].value[value_len] = '\0';
        parser->flags[parser->count].is_boolean = false;
      } else { // It's a boolean flag
        strncpy(parser->flags[parser->count].key, argv[i] + 1, MAX_KEY_LEN - 1);
        parser->flags[parser->count].key[MAX_KEY_LEN - 1] = '\0';
        parser->flags[parser->count].value[0] = '\0'; // No value
        parser->flags[parser->count].is_boolean = true;
      }
      parser->count++;
    }
  }
}

void print_flags(const Parser *parser) {
  for (int i = 0; i < parser->count; i++) {
    if (parser->flags[i].is_boolean) {
      printf("Flag: %s (boolean)\n", parser->flags[i].key);
    } else {
      printf("Flag: %s = %s\n", parser->flags[i].key, parser->flags[i].value);
    }
  }
}

int main(int argc, char *argv[]) {
  Parser parser;
  parse_flags(&parser, argc, argv);
  print_flags(&parser);
  return 0;
}
