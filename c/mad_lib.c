#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *createMadLibs(char *noun, char *verb, char *adjective, char *adverb);

int main(void) {
  char noun[10] = {};
  char verb[10] = {};
  char adjective[10] = {};
  char adverb[10] = {};

  printf("Enter a noun: ");
  scanf("%s", noun);

  printf("Enter a verb: ");
  scanf("%s", verb);

  printf("Enter an adjective: ");
  scanf("%s", adjective);

  printf("Enter an adverb: ");
  scanf("%s", adverb);

  char *result = createMadLibs(noun, verb, adjective, adverb);
  printf("%s\n", result);
  free(result); // free the memory

  return EXIT_SUCCESS;
}

char *createMadLibs(char *noun, char *verb, char *adjective, char *adverb) {
  // Calculate the length of all parameters passed
  const int max_size =
      strlen(noun) + strlen(verb) + strlen(adjective) + strlen(adverb);

  char *value;
  // Allocate a memory to the result
  value = malloc(sizeof(char) * max_size);

  // Concatenate the strings/(char[])
  sprintf(value, "Do you %s your %s %s %s", verb, adjective, noun, adverb);

  return value;
}
