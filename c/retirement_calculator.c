#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main(void) {
  int current_age = 0;
  int retire_age = 0;

  printf("What is your current age? ");
  scanf("%d", &current_age);

  printf("At what age would you like to retire? ");
  scanf("%d", &retire_age);

  // quit program if user is already retire
  if (current_age >= retire_age) {
    printf("You are already retired!!\n");
    return EXIT_SUCCESS;
  }

  time_t current_time = time(NULL);
  char current_year[4];
  strftime(current_year, sizeof(current_year), "%Y", localtime(&current_time));

  int conv = atoi(current_year);
  printf("You have %d years left until you can retire.\n",
         retire_age - current_age);

  printf("%d\n", conv);
  printf("It's %d, so you can retire in %d.\n", conv,
         conv + (retire_age - current_age));
}
