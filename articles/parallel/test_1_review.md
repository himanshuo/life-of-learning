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
High-throughput computing - using large amounts of computing power for long periods of time

High throughput computing is about running long-running jobs that require lots of computing power. It is related to things like batch processing.

You can compare high throughput computing to high performance computing. High Performance computing requires large amounts of computing for short period of time. This is to do large problems quickly.


Examples of high throughput computing are batch processing, parameter space sweep, embarrassingly parallel problems.

Parameter space sweep is literally just running a program multiple times with varying parameters to get a feel for the output.

Embarrassingly parallel is when you split up a problem by simply repeating it with multiple inputs. Little or no effort is needed to separate the problem into a number of parallel tasks. This occurs because there is little or no dependency between the parallel tasks so little communication is required between them.

### Basic model
### load balancing
  you want to run tasks in a balanced way else your performance is limited by slowest machine
  one way to load balance is to make sure that execution time for each task is same
    NOT the same thing as same each has same load. You want to change around load based on how powerful a give processor is.

### queuing systems
### Shared memory
### pThreads
Posix Threads is an execution model where each flow of work is called a thread and the creation and control over these flows is achieved using the POSIX Threads API.

4 groups of Pthread functions
* thread management - creating, joining thread
* mutexes
* condition variables
* synchonization between threads using read/write locks and barriers

????PUT IN SOME EXAMPLE USAGES?????

### synchronization
synchronization of processes is when multiple processes join up at a certain point to aggregate their information and commit to a certain sequence of actions


### races and deadlock
race conditions are when two or more threads can access shared data and they try to change it at the same time

deadlock is when different threads are waiting on each other to finish.

A deadlock situation can arise if all of the following conditions hold simultaneously in a system:

* Mutual exclusion: at least one resource must be held in a non-shareable mode.[1] Only one process can use the resource at any given instant of time.
* Hold and wait or resource holding: a process is currently holding at least one resource and requesting additional resources which are being held by other processes.
* No preemption: a resource can be released only voluntarily by the process holding it.
Circular wait: a process must be waiting for a resource which is being held by another process, which in turn is waiting for the first * process to release the resource. In general, there is a set of waiting processes, P = {P1, P2, …, PN}, such that P1 is waiting for a resource held by P2, P2 is waiting for a resource held by P3 and so on until PN is waiting for a resource held by P1.[1][7]



mutexes for avoiding writing to same memory and thus avoiding race conditions

semaphores for touching memory in correct order

### dance hall architectures
uniformly slow, but scalable architecture.

NUMA machine.

you have processors all connected to interconnection network. Interconnection network then connects to memory.

### interconnection networks
Interconnection networks are used to route data when
* processors need to access memory
* processor needs to share data with another

On one extreme is a shared bus interconnection network. A shared bus is when only a single bus exists for all processors to access memory. This makes it so that only a single processor can access memory at a time.

on the other end is a crossbar switch, where each processor has a unique access path to each memory.

For the case that processors share data with each other, a complete interconnection network contains a path from each processor to every other processor.

There are different shapes for interconnection networks, including ring, mesh interconnection network, hypercube, barrel shifter...

multistage interconnection networks

an interconnection network where processing elements are all on one end of the network, while memory elements are on the other end. Switches connect the two types of elements.

network contention

when two or more stations on a network try to access the network media simultaneously


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

omega network is permutation network. there are diff patterns.

shuffle =
exchange =

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


### throughput versus bandwidth versus latency
same concept: work / time
throughput - work is usually performing some task (not just moving bytes around )
bandwidth - work is usually moving bytes

latency: time / work

### consumer/producer problem
consumer/producer problem. Implement using synchronization things. SO, memorize all of them and how to use them.


### Multistage interconnection networks
omega network
  lgk(N)
  depth of network is ...
  logk N layers
  at each layer,


  just because no contention in switch, doesnt mean no contention. You can have contention at memory. Thus crossbar has contention.


### software combining
1. determine how many
2. determine num memory units
3. goal


### crossbar complexity
cost of switch = O(n^2) because kxk switch.

you have to know both cost, and number of path

exactly one path from src to dest. Thus if anything is blocked, then contentions.

depth = logk(n)

have logk(n) stages

cost = num_stages * cost of each stage = logk(n) * n/k switches per stage * cost of each switch = logk(n) * n/k * O(k^2)
