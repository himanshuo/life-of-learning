#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>
#include <unistd.h>
#define BYTES_TO_READ 1000
int main(){
  char* file_name;
  FILE *fp;
  char buffer[BYTES_TO_READ];


  //read stdin to get file name
  file_name = (char *)malloc(10240 * sizeof(char));
  scanf("%s",file_name);


  //open file
  fp = fopen(file_name, "r");

  //read file
  fgets(buffer, BYTES_TO_READ, (FILE *) fp);

  //stdout file contents
  printf("%s", buffer); //two words

  //seek file
  // if(! fseek(fp, 0, SEEK_SET) ){
    // printf("failed to seek\n");
  // }

  //check if seek worked
  // fgets(buffer, BYTES_TO_READ, (FILE *)fp);
  // printf("%s\n", buffer); //words

  //close file
  fclose(fp);


  //practice writing to file
  fp = fopen(file_name, "w");
  fputs( "three words", fp );
  fclose(fp);
  //check if correctly written to file
  fp = fopen(file_name, "r");
  fgets(buffer, BYTES_TO_READ, fp);
  printf("%s\n", buffer);//three words

  return 0;
}
