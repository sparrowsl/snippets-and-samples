#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void display_employees(char *[], int);
void remove_employee(char *[], char *, int);
bool find_employee(char *[], char *, int);

char *names[] = {"John Smith", "Jackie Jackson", "Chris Jones", "Amanda Cullen",
                 "Jeremy Goodwin"};

int main(void) {
  int array_length = sizeof(names) / sizeof(names[0]);
  char remove[15];

  display_employees(names, array_length);

  printf("Enter an employee name to remove: ");
  fgets(remove, 15, stdin);
  remove[strcspn(remove, "\n")] = 0;

  /*if (find_employee(names, remove, initial_names_length)) {*/
  /*} else {*/
  /*  puts("No employee name found");*/
  /*}*/
  remove_employee(names, remove, array_length);

  printf("Exist: %d\n", find_employee(names, remove, array_length));

  array_length = sizeof(names) / sizeof(names[0]);

  display_employees(names, array_length);
  return EXIT_SUCCESS;
}

void display_employees(char *employees[], int len) {
  printf("There are %d employees:\n", len);
  for (int i = 0; i < len; i++) {
    printf("- %s\n", employees[i]);
  }
}

void remove_employee(char *array[], char *name, int length) {
  for (int i = 0; i < length; i++) {
    if (!strcmp(array[i], name)) {
      printf("found %s\n", name);
      array[i] = array[i + 1];
    }
  }
}

bool find_employee(char *employees[], char *name, int length) {
  for (int i = 0; i < length; i++) {
    if (!strcmp(employees[i], name)) {
      return true;
    }
  }

  return false;
}
