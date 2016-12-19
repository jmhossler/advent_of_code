#include <stdlib.h>
#include <stdio.h>

int main() {
  FILE *fp = fopen("input","r");
  size_t count = 0;
  int floor = 1;

  char temp;
  fscanf(fp,"%c",&temp);
  while(temp != EOF) {
    if(temp == '(') {
      count++;
    }
    else if(temp == ')') {
      count--;
    }
    if(count == -1) {
      printf("Floor: %d\n",floor);
      break;
    }
    ++floor;
    fscanf(fp,"%c",&temp);
    //printf("The count is %lz\n",count);
  }

  //printf("The count is %z\n",count);

  fclose(fp);

  return 0;
}
