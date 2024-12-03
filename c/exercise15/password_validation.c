#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
  char password[20];

  printf("What is the password? ");
  scanf("%s", password);

  if (!strcmp(password, "abc$123")) {
    puts("Welcome!");
  } else {
    puts("I dont' know you.");
  }

  return EXIT_SUCCESS;
}
