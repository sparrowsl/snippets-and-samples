#include <stdio.h>
#include <stdlib.h>

int main(void) {
  // A B C a b c 0 1 2 $ * + /
  printf("A -> %d\n", 'A');
  printf("B -> %d\n", 'B');
  printf("C -> %d\n", 'C');
  printf("a -> %d\n", 'a');
  printf("b -> %d\n", 'b');
  printf("c -> %d\n", 'c');
  printf("0 -> %d\n", '0');
  printf("1 -> %d\n", '1');
  printf("2 -> %d\n", '2');
  printf("$ -> %d\n", '$');
  printf("* -> %d\n", '*');
  printf("+ -> %d\n", '+');
  printf("/ -> %d\n", '/');
  printf("space -> %d\n", ' ');

  return EXIT_SUCCESS;
}
