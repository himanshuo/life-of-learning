Lecture 7
==========
### Readings to do
* pthread
* 2 chapters in book
* paper on consistency symmantics **
* paper on type architectures     **
* readings on hot spots is key to do - will be on midterm, dont need to understand math **
* reading distributed symmetric memory (DSM)
* important to do these readings by thursday since that is when hardware stuff starts
* speedup discussions not covered in class, but will be tested. all from a given paper - ch 7 quinn


### Process vs Thread
Process is basic unit of work in timesharing system. Program in execution.
* stack
* OS resources (open file table)
* Accounting table

Creating process is slow due to
* lots of memory operations
* go between user/kernel

Working with processes is slow because processes have no shared memory and so to make them communicate, you have to go down to kernel level

A thread is within a thread. It has
* stack frame
* PC
* own state
* shared data space
* shared

There are multiple kernel and user threads.

The key reason you want kernels to support threads is that if a user thread enters the kernel, it can block the entire process for things like IO. If kernel knows about these things, the other threads can keep going.

### Thread pitfalls
threads share stack. thus you can overwrite another stack


### message passing versus shared memory
You have to do shared memory synchronization yourself.

If you do message passing, kernel does synchronization for you.

### shared memory synchronization method
we need to control the order that events occur. This is different than the mutual exclusion problem where you simply don't want two things to accessing information at the same time. For us, **ORDER MATTERS**.

you can use semaphores or spin locks to do this.

spin locks uses up too much CPU time so don't use this. Spin locks are for when there is a guarantee that the wait time is very small.

**For the hw, use a barrier data type!!!!!!**

### Semaphore
### locks
usually spin locks.

modern OS's have a test_and_set hardware instruction that makes creating spin locks very easy and fast.

to implement locks, you can just lock the bus. This makes for the test_and_set instruction.
### Barrier
data structure.

initialize it with some value n. There are n threads. None of the n threads are allowed to go forward until all of them are able to come to the barrier location.

### Event
similar to semaphore in that you wait on it.

### Deadlocks
avoid them by
* avoid circular waits - ask for things in particular order
* have a list of things to wait for. Ask for them, or release them all at once
