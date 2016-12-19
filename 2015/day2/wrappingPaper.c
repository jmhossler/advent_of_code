#include <stdio.h>
#include "scanner.h"

int min3(int, int, int);
int min2(int, int);
int squareFeet(int, int, int);
int mult(int, int, int);
int ribbonFeet(int, int, int);

int main() {
  int x, y, z, total = 0;
  FILE *fp = fopen("input","r");

  x = readInt(fp);
  readChar(fp);
  y = readInt(fp);
  readChar(fp);
  z = readInt(fp);

  while(!feof(fp)) {
    total += ribbonFeet(x,y,z);
    x = readInt(fp);
    readChar(fp);
    y = readInt(fp);
    readChar(fp);
    z = readInt(fp);
  }

  printf("Total: %d\n",total);

  return 0;
}

int ribbonFeet(int w, int h, int l) {
  int feet = mult(w,h,l);
  int x = min3(w,h,l);
  int y = 0;
  if(x == w) {
    y = min2(h,l);
  }
  else if(x == h) {
    y = min2(w,l);
  }
  else {
    y = min2(w,h);
  }

  feet += 2 * x + 2 * y;

  return feet;
}

int squareFeet(int w, int h, int l) {
  int x,y,z;
  x = l * w;
  y = w * h;
  z = h * l;
  int sqft = 2*x + 2*y + 2*z + min3(x,y,z);

  return sqft;
}

int min3(int x, int y, int z) {
  if(x <= y && x <= z) {
    return x;
  }
  else if(y <= x && y <= z) {
    return y;
  }
  else if(z <= y && z <= x) {
    return z;
  }
  return 0;
}

int mult(int x, int y, int z) {
  return x * y * z;
}

int min2(int x, int y) {
  if(x <= y) { return x; }
  else { return y; }
}
