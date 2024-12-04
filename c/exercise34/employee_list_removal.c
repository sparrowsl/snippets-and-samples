#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void remove_employee(char *, char *);

char *names[] = {"John Smith", "Jackie Jackson", "Chris Jones", "Amanda Cullen",
                 "Jeremy Goodwin"};

int main(void) {
  char remove[15];

  int names_length = (int)(sizeof(names) / sizeof(names[0]));

  printf("There are %d employees:\n", names_length);
  for (int i = 0; i < names_length; i++) {
    printf("- %s\n", names[i]);
  }

  printf("Enter an employee name to remove: ");
  scanf("%s", remove);

  return EXIT_SUCCESS;
}

void remove_employee(char *array, char *name) {
  int names_length = (int)(sizeof(names) / sizeof(names[0]));

  for (int i = 0; i < names_length; i++) {
  }
}
