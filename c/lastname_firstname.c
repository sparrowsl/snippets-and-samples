#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
  char name[50];
  char *firstname;
  char *lastname;

  printf("Enter a first name and last name: ");
  fgets(name, 50, stdin);
  name[strcspn(name, "\r\n")] = 0;

  char *token = strtok(name, " ");
  firstname = token;
  while (token != NULL) {
    lastname = token;
    token = strtok(NULL, " ");
  }

  printf("%s %c.\n", lastname, firstname[0]);

  return EXIT_SUCCESS;
}
