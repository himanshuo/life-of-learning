Lecture 8 - Process Synchronization
=============

### race conditions
A race condition when separate processes manipulate the same information at the same time and so it becomes difficult to determine what the actual value of the information should be.

There are 2 reasons why a race condition occurs
* different processes access shared resources
* assembly code is interleaved

### Example revealing shared resource usage race condition
Note that this is actually from end of lecture 7, but yea, its a totally new topic so i put it here.
    Process 1:
      while(1){
        disable();
        x++;
        enable();
      }

    Process 2:
      while(1){
        disable();
        x--;
        enable();
      }

This leads to a race condition between process 1 and process 2 since we do not know what x is at a given time.


#### Example revealing assembly code interleaving leading to race condition (also in bottom of lecture 1)
Concurrency is basically making multiple things happen at the same time. Interweaving processes at the assembly level can cause *race conditions*, where something is edited twice and the changes aren't synchronized correctly.

![](overall_idea_of_OS/ae8a5fbbd529d0db9f19d2e966265227.png)

Process 1 increments count, while process 2 decrements count.
The assembly for process 1 would look like
* (1) load: r1 <- count
* (2) increment: r1 <- r1+1
* (3) store: count <- r1
The assembly for process 2 would look like
* (1') load: r2 <- count
* (2') decrement: r2 <- r2-1
* (3') store: count <- r2

Assume count is initially 3. If P1 ran then P2 ran, count would be 3. If P2 ran then P1 ran, count would be 3.

Instruction level concurrency is about running the assembly instructions concurrently so they still result in the same answer. This is a difficult thing to do. The above would not work with instruction level concurrency. An example that shows that you cannot just interweave the assembly instructions above is:

(1) (2) (1') (2') (3) (3') which is:
* (1) load: r1 <- count       ; count=3, r1=3
* (2) increment: r1 <- r1+1   ; count=3, r1=4
* (1') load: r2 <- count      ; count=3, r1=4, r2=3
* (2') decrement: r2 <- r2-1  ; count=3, r1=4, r2=2
* (3) store: count <- r1      ; count=4, r1=4, r2=2
* (3') store: count <- r2     ; count=2 , r1=4, r2=2

Thus this interleaving would make count=2 when count should be 3.






two privileged instructions

disable()
enable()




2 major problems
* malicious user can have control of cpu for ever
* solution doesnt work on multiprocessors
