#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void display_employees(char *[], int);
void remove_employee(char *[], char *, int *);
bool find_employee(char *[], char *, int);

int main(void) {
  FILE *file = fopen("./employees.txt", "r");
  if (file == NULL) {
    perror("Error opening file");
    return 1;
  }

  char *names[5]; // array of pointers to strings
  char buffer[100];

  int i = 0;
  while (i < 5 && fgets(buffer, sizeof(buffer), file)) {
    buffer[strcspn(buffer, "\n")] = '\0';
    names[i] = strdup(buffer); // dynamically allocate and copy the string
    i++;
  }
  fclose(file);

  // never recalculate the array length after first calculation.
  int array_length = sizeof(names) / sizeof(names[0]);
  char name_to_remove[15];

  display_employees(names, array_length);

  printf("Enter an employee name to remove: ");
  fgets(name_to_remove, 15, stdin);
  name_to_remove[strcspn(name_to_remove, "\n")] = 0;

  remove_employee(names, name_to_remove, &array_length);

  display_employees(names, array_length);
  for (int i = 0; i < array_length; i++) {
    free(names[i]);
  }
  return EXIT_SUCCESS;
}

void display_employees(char *employees[], int len) {
  printf("There are %d employees:\n", len);
  for (int i = 0; i < len; i++) {
    printf("- %s\n", employees[i]);
  }
}

void remove_employee(char *array[], char *name, int *length) {
  for (int i = 0; i < *length; i++) {
    if (!strcmp(array[i], name)) {
      printf("found %s\n", name);
      // Shift elements to the left
      for (int j = i; j < *length - 1; j++) {
        array[j] = array[j + 1];
      }

      (*length)--;
      return;
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
