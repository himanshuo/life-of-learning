#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>


int main(){
  int * int_ptr = malloc(sizeof(int));
  int * origin = int_ptr;

  //original
  *int_ptr = 0;
  display(origin, 1, "original");

  //pointer++
  int_ptr++; //increments pointer by sizeof(T)
  display(origin, 2, "pointer++");

  //*pointer++
  *int_ptr++; //increments pointer by sizeof(T)
  //thus what is happening is that by default, the order of operations are: *(int_ptr++)
  // ++ happens before dereferencing 
  display(origin, 3, "*pointer++");

  //(*pointer)++
  (*int_ptr)++; //increments pointer by sizeof(T)
  display(origin, 3, "(*pointer)++");



  return 0;
}


void display(int * ptr, int n, char str[] ){
  int i=0;
  printf("-----------%s------------\n", str);
  for(i=0;i<n;i++){
    printf("%p -> %d\n", ptr + i, *(ptr+i));
  }

}
