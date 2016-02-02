Class 4 Terms  - Shared Memory Computing
==========================================

### Pthreads
Posix Threads is an execution model where each flow of work is called a thread and the creation and control over these flows is achieved using the POSIX Threads API.

Basically, these are just threads and Pthreads is just the Linux version of them.

All the procedures with pthreads are prefixed with "pthread_". They can be categorized into 4 groups:
* thread management - creating, joining thread
* mutexes
* condition variables
* synchonization between threads using read/write locks and barriers

### OpenMP
API that supposed multi-platform shared memory multiprocessing. It is for c, c++.

OpenMP uses a portable, scalable model that gives programmers a simple and flexible interface for developing parallel applications.

### race conditions
when two or more threads can access shared data and they try to change it at the asme time.
### synchonization
synchronization of processes is when multiple processes join up at a certain point to aggregate their information and commit to a certain sequence of actions.

### locks / mutex
a thread acquires a lock before performing its work and then releases it when it finishes.
Locks are a general term. Spin locks, and a binary semaphore are both locks.

### spin locks
a thread simply waits ("spins") until the lock becomes available.

### semaphores
variable used to determine how many units of a particular resource is available. If there are more variables available, then a new process that requires the resource can use it and decrement the unit count. Once a process finishes using the resource, the unit count is incremented. If a process wants the resource and the unit count is 0, then the process must wait for some other thread to release its resource.
### interconnection networks
Interconnection networks are used to route data when
* processors need to access memory
* processor needs to share data with another

On one extreme is a *shared bus* interconnection network. A shared bus is when only a single bus exists for all processors to access memory. This makes it so that only a single processor can access memory at a time.  

on the other end is a crossbar switch, where each processor has a unique access path to each memory.

For the case that processors share data with each other, a complete interconnection network contains a path from each processor to every other processor.  

There are different shapes for interconnection networks, including ring, mesh interconnection network, hypercube, barrel shifter...

### multistage interconnection networks
an interconnection network where processing elements are all on one end of the network, while memory elements are on the other end. Switches connect the two types of elements.
### network contention
when two or more stations on a network try to access the network media simultaneously
### bandwidth
rate of data transfer
### latency
the delay before a transfer of data begins following an instruction for its transfer
### memory hierarchy
![](class_4_terms/f09a84fd32110b41568ba2d5a0fc7fa4.png)
