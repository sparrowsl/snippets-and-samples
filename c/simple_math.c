#include <stdio.h>
#include <stdlib.h>

float add(float a, float b);
float subtract(float a, float b);
float multiply(float a, float b);
float divide(float a, float b);

int main(void) {
  float first_number = 0;
  float second_number = 0;

  printf("What is the first number? ");
  scanf("%f", &first_number);

  printf("What is the second number? ");
  scanf("%f", &second_number);

  float sum = add(first_number, second_number);
  float difference = subtract(first_number, second_number);
  float product = multiply(first_number, second_number);
  float quotient = divide(first_number, second_number);

  printf("%.2f + %.2f = %.2f\n", first_number, second_number, sum);
  printf("%.2f - %.2f = %.2f\n", first_number, second_number, difference);
  printf("%.2f * %.2f = %.2f\n", first_number, second_number, product);
  printf("%.2f / %.2f = %.2f\n", first_number, second_number, quotient);

  return EXIT_SUCCESS;
}

float add(float a, float b) { return a + b; }

float subtract(float a, float b) { return a - b; }

float multiply(float a, float b) { return a * b; }

float divide(float a, float b) { return b == 0 ? 0 : (a / b); }
