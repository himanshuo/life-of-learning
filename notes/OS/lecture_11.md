Lecture 11
=============

Monitors
*
*

bounded-buffer (multiple producers, multiple consumers)
monitor bounded-buffer
  * global varibales initialized in bounded buffer

    char buffer[n];
    int nextin = 0, nextout = 0, full_count=0; //these have to be gaurded  
    Condition notempty, notfull;



    deposit(char c){
        if(full_count == n) notfull.wait; // waits until not full
        buffer[nextin] = c;
        nextin = (nextint) % n;
        full_count++;
        notempty.signal; // last process  so do not need to block
    }


    remove(char c){
      if(full_count==0) notempty.wait;
      c = buffer[nextout];
      nextout = (nextout) % n;
      full_count--;
      notfull.signal;
    }


cobegin
  P1 // P2 // P3 // P4 // P5
coend


P1:
  ........
  deposit(c);
  .......

C1:
  ........
  remove(c);
  ........

    //reader-writer problem is when multiple readers can access resource that is written on by multiple writers
    monitor Reader_writers {
      int read_cnt=0, writing=0;
      condition ok_to_read, ok_to_write;
    }


    start_read(){
      if(writing){
        ok_to_read.wait;
      }
      read_cnt++;
      ok_to_read.signal;
   }


   finish_read(){
     read_cnt--;
     if(read_cnt==0){
       ok_to_write.signal;
     }
   }


   start_write(){
     //either someone is reading file or writing to file
     if(read_cnt!=0 || writing){
        ok_to_write.wait;
     }
     writing=1;
   }

   finish_write(){
     writing=0;
     if(!empty((ok_to_read)){
       ok_to_read.signal;
     } else {
       ok_to_write.signal;
     }
   }


reader:
  ....
  start_read()
  file access
  finish_read()
  ....





### How to simulate the 3 things using semaphores
* external functions for the entry points
  * must enforce mutual exclusion
* conditional variables
* signal queue

##### External functions for the entry points

    bp(mutex);
    .....
    body of F;
    .....
    if(next_count >0){
      v(next); //next is a semaphore used by signal processes
    } else {
      bV(mutex);
    }

##### Condition variables using semaphore
    //x_count was initialized to 0
    //x_count
    x_count++;
    if(next_count > 0){
      v(next);//release one of the processes. Brought back into monitor
    } else {
      bV(mutex);
    }
    P(x_sem);
    x_count--;

##### signal queue using semaphore
    if(x_count > 0){
      next_count++;
      v(x_sem);
      p(next);
      next_count--;
    }



### Producer-consumer problem in C
pthread_cond_wait, pthread_cond_signal is needed for project 2

    cond_t empty, full;
    mutex_t mutex; //always initialized to 1 (free)

    void * producer(void * arg){
      item_t i;
      while(1){
        pthread_mutex_lock(&mutex);               //P1
        while(count == max){ //count is number of things put into consumer producer problem                //P2
          pthread_cond_wait(&empty, &mutex);               //P3
        }
        put(i);                                          //P4
        pthread_cond_signal(&full);                        //P5
        pthread_mutex_unlock(&mutex);                 //P6
      }
    }



    void * consumer(void * arg){
      item_t i;
      while(1){       
        pthread_mutex_lock(&mutex);     //C1
        //while instead of if because   
        while(count==0){                //C2
          pthread_cond_wait(&full, &mutex); //C3
        }
        int temp=get();                 //C4
        pthread_cond_signal(&empty);    //C5
        pthread_mutex_unlock(&mutex);   //C6
      }
    }

    int buffer[MAX];
    int fill = 0;
    int use=0;
    int count=0;
    void put(int value){
      buffer[fill] = value;
      fill = (fill+1)%MAX;
      count++;
    }
    int_t get(){
      int_t temp = buffer[use];
      use = (use+1) % max;
      count--;
      return temp;
    }






    Max=1
    consumer 1(c1, c2,c3) -> producer 1(P1,P2, p3, p4, p5, p6, 1, p2, p3)

    -> consumer2(c1,c2,c4,c5,c6)
    -> consumer1(c4)


# reusable resources
DEADLOCK

    P1:
      ....
      open(f1,w);
      open(f2,w);
      ....

    P2:
      ....
      open(f2,w);
      open(f1,w);
      ....


# consumable resources 
