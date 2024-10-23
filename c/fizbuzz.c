// FizzBuzz program that prints numbers from 1 to 100,
// but replaces multiples of 3 with "Fizz"
// and multiples of 5 with "Buzz".
#include <stdio.h>
#include <stdlib.h>

int main(void) {

  for (int i = 1; i <= 100; i++) {
    if (i % 3 == 0) {
      puts("Fizz");
    } else if (i % 5 == 0) {
      puts("Buzz");
    } else {
      printf("%d\n", i);
    }
  }

  return EXIT_SUCCESS;
}
