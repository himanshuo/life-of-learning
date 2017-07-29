Lecture 11 - Review Session
============

##### Exam
The exam will be closed book, closed note, closed neighbor, closed electronic devices. It will be 75 minutes long. There will be 5-7 questions on the exam. There will be simple coding required. Any function prototypes you need will be provided. Do not waste time remembering the function prototypes for the pthreads library functions. Make sure you know what they do!

Points back for leaving question blank. You can cross it all off and write NO BS on it to get points back. 

So far we have covered:
         Basic intro material.
         Motivation for parallel computing,
         basic hardware models,
         speed-up & efficiency,
         granularity,
         scalability,
         functional and data parallelism,
         Fosters methodology.
         Sequential program optimization:
           High throughput computing:
            Basic model,
            load balancing,
            queuing systems.
           Shared memory.
           pThreads,
           synchronization,
           races and deadlock,
           dance hall architectures,
           simple analysis of multistage interconnection networks,
           hot spots,
           combining,
           reduction operations,
           cache and cache coherency,
           snoopy caches,
           directory–based caches.

Exam questions may come out of the homework as well.
For example, questions using pthreads and code optimization.

The slides are all available on collab as well.

In terms or reading:
Pacheco chapters: 1, 2, 4
Quinn chapters: 7, 14 (optional 1-2 – similar overview fluff as in 1-2 of Pacheco)

Papers:
Type Architectures
High throughput computing in the sciences
Consistency Semantics
Misleading performance …
The three hot spot papers


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



UNDERSTAND ALIGNMENT OF VARIABLES IN C.

*CACHE ALIGNMENT*
  a cache has rows and 2 columns
  each row is related.
  column 1: tag
  column 2: cache line ()

  If a data structure is 'cache aligned', then reading entire word can be done cache.
  This will only happen if the entire word is stored in a given cache line
  This will require that the data structure has size = cache line

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


*cache blocking*
for i: 0->10000
  for j: 0->10000
    // load x[i] into cache
    x[i][j] = x[i,j] * y[j,i]



LOADBALANCING
  you want to run tasks in a balanced way else your performance is limited by slowest machine
  one way to load balance is to make sure that execution time for each task is same
    NOT the same thing as same each has same load. You want to change around load based on how powerful a give processor is.


same concept: work / time
throughput - work is usually performing some task (not just moving bytes around )
bandwidth - work is usually moving bytes

latency: time / work


consumer/producer problem. Implement using synchronization things. SO, memorize all of them and how to use them.
