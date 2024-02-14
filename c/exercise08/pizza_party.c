#include <stdio.h>
#include <stdlib.h>

int main(void) {
  int people = 0;

  printf("How many people? ");
  scanf("%d", &people);

  if (people <= 0) {
    printf("Please enter a valid number!!\n");
    exit(1);
  }

  int pizzas = 0;
  printf("How many pizzas do you have? ");
  scanf("%d", &pizzas);
  if (pizzas <= 0) {
    printf("Please enter a valid number!!\n");
    exit(1);
  }

  int leftovers = people % pizzas;
  printf("%d people with %d pizzas\nEach person gets %d pieces of pizza.\n",
         people, pizzas, (people / pizzas) / 2);
  printf("There are %d leftovers pieces.\n", leftovers);
}
