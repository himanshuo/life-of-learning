Why Parallel Computing?
======================

# We need ever-increasing performance
A few problems that have become possible due to increased performance are
* Climate modeling
* Protein folding
* Drug discovery
* energy research
* data analysis

# We need Parallel systems

Industry is making transistors more and more dense. They are putting multiple complete processors onto a single chip in order to use each processor in parallel. Industry feels this is the next step in making chips faster. Making them smaller has been sooo overdone at this point that its hard to make them even smaller. Thus, integrated circuits now are *multicore* processors. A *core* is now synonymous with a central processing unit (CPU). Thus a single CPU is called a single-core system.

# We need to write Parallel Programs

It is difficult to write programs that automatically convert serial programs into parallel programs. The reason for this is that the constructs used in serial programs do not always lend themselves to parallel programming. To turn a serial program into a parallel program, you usually have to understand the input/output of the serial program and determine a new algorithm entirely that has the same input/output but does so in a parallel way.

As an example, think of a program that sums a bunch of numbers together. Going from serial to parallel, the obvious change in the code should be to make each core calculate a portion of the sum and then some last core sums all the sums to get the final result. This is faster than normal serial however, not too much faster since the final core has to still sum up a lot of numbers. A smarter way to do this would be to have the cores sum neighboring numbers and then neighboring cores would sum their values. Thus you'd have a tree like structure of cores each summing the result of the previous height of the tree. Of course, now you have a completely different algorithm.


# How to write parallel programs

You have to partition the work to be done among different cores.
## task-parallelism
various tasks are partitioned among cores
## data-parallelism
data is partitioned among cored


For example, a teacher can get a bunch of TA's to grade papers in parallel.
If the teacher decides to do task-parallelism, then each grader would basically just continue working and as soon as he finishes grading a new paper, the teach would hand him another until all the papers are graded. Thus slow graders would grade less papers, fast would grade more. If the teacher decides to do data-parallelism, then each grade would be assigned x papers and each grader would only have to grade until he finishes his work. In both cases, graders are working in concurrently.

## communication
Each of the cores needs to communicate with the other to effectively do the work in parallel. Each core needs to do a bit of work and then send its results to another core or at least have some way among the cores to handle the results of each individual core.

## load balancing
Cores that handle the results of other cores cannot be overworked. You generally want each 'handler' core to have equal amounts of work. Thus a load balancer matches results of individual cores to 'handler' cores.

## synchronization
Until action x is done, do not do action y.

The idea is that parallel programs still have to make sure that some things are in fact done synchronously because of simply their nature. For example, a program that sums a list of number that the user provides MUST first determine the list of numbers from the user and THEN sum them up.

# Key Concepts

## Shared Memory versus Distributed memory
In shared memory systems, cores share access to computer's memory. Each core can read and write each bit of memory. The cores coordinate with each other by having them examine and update shared-memory locations. Pthreads and OpenMP are for shared memory systems. OpenMP is more high level and thus faster to write than Pthreads.

In distributed memory systems, each core has it's own private memory and must communicate explicitly with other cores. Message Passing Interface (MPI) is for distributed systems

## Concurrent, Parallel, Distributed

Concurrent = program in which multiple tasks can be in progress at any instant  
Parallel = program in which multiple tasks cooperate closely to solve a problem  
Distribute = program which needs to cooperate with other programs to solve a program

Thus parallel and distributed programs are concurrent.

Distributed programs tend to be more "loosely coupled"
