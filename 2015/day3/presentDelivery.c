#include <stdlib.h>
#include <stdio.h>
#include "scanner.h"

int size;

int **initArray();
int **resize(int **);
int countHouses(int **);

int main() {
  size = 10000;

  int x1 = size/2, y1 = size/2, x2 = size/2, y2 = size/2;
  printf("Here\n");
  int **houses = initArray();
  printf("Here2\n");
  houses[x1][y1] = 1;

  FILE *fp = fopen("input","r");

  char dir = readChar(fp);
  char dir2 = readChar(fp);

  while(!feof(fp)) {
    if(dir == '^') {
      --y1;
    }
    else if(dir == 'v') {
      ++y1;
    }
    else if(dir == '<') {
      --x1;
    }
    else if(dir == '>') {
      ++x1;
    }
    if(dir2 == '^') {
      --y2;
    }
    else if(dir2 == 'v') {
      ++y2;
    }
    else if(dir2 == '<') {
      --x2;
    }
    else if(dir2 == '>') {
      ++x2;
    }

    houses[x1][y1] = 1;
    houses[x2][y2] = 1;
    dir = readChar(fp);
    dir2 = readChar(fp);
  }

  int count = countHouses(houses);

  printf("Houses: %d\n",count);

  return 0;
}

int countHouses(int **arr) {
  int count = 0;
  int i,j;
  for(i = 0; i < size; ++i) {
    for(j = 0; j < size; ++j) {
      if(arr[i][j] == 1) {
        count++;
      }
    }
  }
  return count;
}

int **initArray() {
  int **array = malloc(sizeof(int *) * size);

  int i,j;
  for(i = 0; i < size; ++i) {
    array[i] = malloc(sizeof(int) * size);
    for(j = 0; j < size; ++j) {
      array[i][j] = 0;
    }
  }

  return array;
}

int **resize(int **arr) {
  int s = size, i, j;
  size *= 2;
  int **temp = initArray();
  for(i = 0; i < s; ++i) {
    for(j = 0; j < s; ++i) {
      temp[i][j] = arr[i][j];
    }
  }
  return temp;
}
