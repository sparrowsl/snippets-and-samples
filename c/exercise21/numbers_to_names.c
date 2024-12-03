#include <stdio.h>
#include <stdlib.h>

char *get_month(int);

int main(void) {
  int month_idx = 0;

  printf("Please enter the number of the month: ");
  scanf("%d", &month_idx);

  if (month_idx < 1 || month_idx > 12) {
    puts("No month found!!");
  } else {
    char *month_name = get_month(month_idx);
    printf("The name of the month is %s\n", month_name);
  }

  return EXIT_SUCCESS;
}

char *get_month(int idx) {
  switch (idx) {
  case 1:
    return "January";
  case 2:
    return "February";
  case 3:
    return "March";
  case 4:
    return "April";
  case 5:
    return "May";
  case 6:
    return "June";
  case 7:
    return "July";
  case 8:
    return "August";
  case 9:
    return "September";
  case 10:
    return "October";
  case 11:
    return "November";
  case 12:
    return "December";
  default:
    return "No month found!!";
  }
}
