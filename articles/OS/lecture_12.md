Lecture 12
============

P1:
  ...
  if(c) send(p2, m);
  while(1){
    receive(p3,m);
    ...
    send(p2,m);
  }

P2:
  ...
  while(1){
    receive(p1,m);
    ...
    send(p3,m);
  }

P2:
  ...
  while(1){
    receive(p2,m);
    ...
    send(p1,m);
  }


deadlock prevention
  1. mutual exclusion must be enforced
  2. no preemption must be enforced
  3. hold-and-wait must be enforced
  4. circular wait must hold




Deadlock prevention schemes
* new deadlock prevention
* wount-wait



Prioty
Pi, Pj

P(pi) > P(pj)
11        11
5         10
     <



P1 -> P2 -> P3

wait-die
  timestamp
  Pi,    Pj
  11     11
 TS(5)  TS(10)
wound-wat
Pi, Pj




deadlock avoidance
  safe state 1 -> safe state 2



if system in safe state -> all running processes can be satisfied ...
never go from safe state to unsafe state   



12 tape drives
|process|  max needs    |  allocate |   still needed|
|---|---|---|---|
|p0|10|5|5|
|p1|4|2|2|
|p2|9|2|7|

he summed up the allocated column and wrote 9 underneath it.

12-9=3 (circled 3) (crossed off 3 and replaced with 2)

he crossed off 2 in allocated, 9 in the sum and replaced it with 10, 7 in still needed and replaced it with 6.




Deadlock detection
(graph reduction method)


request allocation graph



deadlock recovery
* priority of process
  * realtime, iteractive, batch
* cost of restarting the process
  * shell, internet browser, text editors, finite element software
* the current state of process


SKIP ch6.
going to do ch7.

ch 7 - memory management


logical-to-physical address binding/assignment

static binding
* earliest moment
  low kernel code, real time, embedded
* compilation
* linking time

dynamic binding
* relocatable


relocatable:     load R1, tmp
absolute: R1, immediate operands
