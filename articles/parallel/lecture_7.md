Lecture 7
==========
### Readings to do
* pthread
* 2 chapters in book
* paper on consistency symmantics
* paper on type architectures
* readings folder is basically for everyone
* readings on hot spots is key to do - will be on midterm, dont need to understand math
* reading distributed symmetric memory (DSM)
* important to do these readings by thursday since that is when hardware stuff starts
* speedup discussions not covered in class, but will be tested. all from a given paper - lecture 7 quinn



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
