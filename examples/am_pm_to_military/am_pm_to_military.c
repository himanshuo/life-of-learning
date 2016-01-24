#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>
# define NUMTESTCASES 12
void time_convert(char * time);


int main(){
    // char* time = (char *)malloc(10240 * sizeof(char));
    // scanf("%s",time);

    int i;


    char testCases[NUMTESTCASES][100] = {
      "00:00:00AM",
      "00:00:00PM",
      "12:00:00AM",//00:00:00
      "12:00:00PM",//12:00:00
      "03:00:00AM",
      "03:00:00PM",
      "07:05:45AM",
      "07:05:45PM",
      "07:15:00PM",
      "07:05:12PM",
      "11:05:56PM",
      "11:45:00PM",
    };

    for(i=0;i<NUMTESTCASES;i++){
        time_convert(testCases[i]);
    }


    return 0;
}


void time_convert(char * time){
  char * hour;
  char * min;
  char * sec;
  char * pm_portion;
  bool PM;
  int hourNum;

  printf("%s->", time);
  hour = strtok(time, ":");
  min = strtok(NULL, ":");
  sec = strtok(NULL, ":");

  pm_portion = strstr(sec, "PM");
  // printf("pm_portion is: %s\n", pm_portion);
  if(pm_portion){
    PM = true;
  } else {
    PM = false;
  }

  sec[2] = '\0';

  hourNum = atoi(hour);

  // printf("PM is: %d\n", PM);
  if(hourNum == 12){
      if(!PM){
        hourNum = 0;
      }
  } else if(PM){
    hourNum = hourNum + 12;
  }

  //result
  printf("%02d:%s:%s\n",hourNum, min, sec);

}
