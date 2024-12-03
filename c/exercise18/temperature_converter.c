#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>

int to_celcius(int);
int to_fahrenheit(int);

int main(void) {
  char choice = '0';
  int temperature = 0;

  puts("Press C to convert from Fahrenheit to Celcius");
  puts("Press F to convert from Celcius to Fahrenheit");
  printf("Your choice: ");
  scanf("%c", &choice);

  if (toupper(choice) == 'C') {
    printf("Please enter the temperature in Fahrenheit: ");
    scanf("%d", &temperature);
    printf("The temperature in Celcius is %d\n", to_celcius(temperature));
  } else if (toupper(choice) == 'F') {
    printf("Please enter the temperature in Celcius: ");
    scanf("%d", &temperature);
    printf("The temperature in Fahrenheit is %d\n", to_fahrenheit(temperature));
  } else {
    puts("Invalid option, choose 'C' or 'F'");
    return EXIT_SUCCESS;
  }

  return EXIT_SUCCESS;
}

int to_celcius(int fahren) {
  int celcius = (fahren - 32) * 5 / 9;
  return celcius;
}

int to_fahrenheit(int celcius) {
  int fahrenheight = (celcius * 9 / 5) + 32;
  return fahrenheight;
}
