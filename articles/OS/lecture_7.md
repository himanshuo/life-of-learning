Lecture 7 - threads
==============


1. PC
2. stack
3. registers

1. specialist paradigm
given task that can be done
2. client server paradigm
client 1 -> server 1
client 2 -> server 3
client 3 -> server 2


word processor has threads:
  * user input
  * spell/grammar checking
  * document layout
  * ...


3. assembly line paradigm
task -> thread 1 -> thread 2 -> thread 3 ->
compiler (2 passes)
pass 1: creates symbol table
pass 2: symbol table
pass 3: object code

many to one (solaris)



int sum; //in global space. inter thread sharing
int main(int argc, char * argv[]){
    pthread_t tid;
    pthread_attr_t attr;
    pthread_attr_init(&attr);
    pthread_create(&tid, &attr, runner, argv[1]);
    pthread_join(tid, NULL);
    printf("sum=%\n", sum);
}



void * runner(void * param){
  int upper = atoi(param);
  sum =0;
  for(int i=0;i<=upper; i++){
    sum += i;
  }
  pthread_exit(0);
}


//creates multiple threads.
//id of each is stored in tid array
pthread_t tid[len];
for(int i=0; i<len; i++){
  pthread_create(&tid[i] ,NULL, runner, NULL);
}
//stop threads
for(int i=0;i<len;i++){
  pthread_join(tid[i], NULL);
}
