Lecture 6
==============

### Boundary value Problem
Imagine a rod with hot water to the left,

rod cools as time progress

at time 0, highest temperature. At time increases, time decreases.

DIAGRAM


Partitioning
* one data item per grid point
* associateone primitive task with each grid oint
* two dimensional doman decomposition

Communication
* identify communication pattern between primitive tasks
* each interior primitive task has three incoming and three outgoing channels


### Reduction Problem

Associative operator: O

a0 O 11 O a2 O a3 O ... O a{n-1}

O can stand for +, -, /, * , ...

performing each O sequentially is slow. You want to speed this up by doing each operation in parallel.  

DIAGRAM 
