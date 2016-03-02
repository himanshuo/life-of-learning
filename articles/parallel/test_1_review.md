Test 1 Review
============

### Motivation for parallel computing
### basic hardware models
### speed-up & efficiency
Speedup = sequential execution time / parallel execution time

Note that speedup is supposed to be > 1 if things are going correctly.

Conventional Speedup is basically just the reduction in execution time. If you run a slow version and fast version of a solution on a slow computer then you get a much greater speedup than if you made the comparison on a fast computer. Thus it makes it seem like your speedup is not as significant when you run it on a fast computer - *normal speedup penalizes faster absolute speed*. This is bad.

*Scaled Speedup* to the rescue! What scaled speedup does is simply scale the speedup equation so that you also consider the speed of the computer when determining speedup.

????EQUATIONS FROM CH IN BOOK ????

### granularity
### scalability
Scalability is the capability of a system to handle growing amounts of work.

Horizontal scaling - add more nodes to a system  

Vertical scaling - add more capability to a single node in the system

Capability computing - maximum computing power to solve a single large problem in the shortest amount of time.

Capacity computing - efficient cost-effective computing power to solve a lots of easy problems.

Super computers generally aim for maximum in capability computing rather than capacity computing.

Amdahl's Law - measure the maximum potential speedup.  

    30% of a program is sped up by half.
    old: (3/10)+(7/10) = 1
    new: (3/10) * (1/2) +(7/10) = 17/20
    speedup = 1 / 17/20 = 20/17 = 1.18

### functional and data parallelism
### Fosters methodology
### Sequential program optimization
### High throughput computing
### Basic model
### load balancing
### queuing systems
### Shared memory
### pThreads
### synchronization
### races and deadlock
### dance hall architectures
### simple analysis of multistage interconnection networks
### hot spots
### combining
### reduction operations
### cache and cache coherency
### snoopy caches
### directory–based caches
### Bit shuffling

Bit shuffling - shared memory ppt, slide 16
 may be asked to show how bits travel through network
 may be asked to let me know when tree saturation occurs.
 packets are injected from start, they go through switches and queues, if blocked then backpropogate, go to
 how does cost function for these work
  shrink depth of tree leads to increases cost
he would give source and destination, we have to draw path
he may ask to give him access pattern that will clog tree
functions of the buffers and the queues: assume no queues, then everything will go to hot sink. switch will
using this pic, say we want to go from 5->3  (101 -> 011).
  from top to bottom, top=0, bottom=7. From column 0, you are at row 5.
    5->4->


### cache alignment
UNDERSTAND ALIGNMENT OF VARIABLES IN C.

*CACHE ALIGNMENT*
  a cache has rows and 2 columns
  each row is related.
  column 1: tag
  column 2: cache line ()

  If a data structure is 'cache aligned', then reading entire word can be done cache.
  This will only happen if the entire word is stored in a given cache line
  This will require that the data structure has size = cache line

### cache line
*CACHE LINE SIZE and how it affects programs*
struct of array versus array of structs

rise and fall of shared memory machines:
* fundamental problem - interconnection networks  
  * dance hall
  * busses dont scale
  * crossbar (know cost and performance) is ideal but it still has contentions
    * "log k n" contention ????
  * other extreme (know cost and performance)
* you end up with NUMA which is bad.
* shared memory systems are easier for programmers to understand
* PRAM cannot be implemented
* at small scale, shared memory machines work.
* WILL BE QUESTIONS ON THIS TOPIC - why did shared memory machines fail for large scale?

### cache blocking
*cache blocking*
for i: 0->10000
  for j: 0->10000
    // load x[i] into cache
    x[i][j] = x[i,j] * y[j,i]


### load balancing
  you want to run tasks in a balanced way else your performance is limited by slowest machine
  one way to load balance is to make sure that execution time for each task is same
    NOT the same thing as same each has same load. You want to change around load based on how powerful a give processor is.


### throughput versus bandwidth versus latency
same concept: work / time
throughput - work is usually performing some task (not just moving bytes around )
bandwidth - work is usually moving bytes

latency: time / work

### consumer/producer problem
consumer/producer problem. Implement using synchronization things. SO, memorize all of them and how to use them.
