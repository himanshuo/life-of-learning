Lecture 2- Things we will Learn
===============================

### Resource Management
resource = anything on the hardware or software  
### What is a Process?
A Process is *program in execution*. This is key. A program MUST be in a execution state in order for the program to be called a process.   
A process is a resource.  
A process can access memory space, registers, files, any of the normal things you think about in a computer.

The execution environment needs to be saved and restored when running processes concurrently.

### State Transition
A process must be able to save it's state. This is necessary because the contents of registers that the process is accessing can be contaminated by other processes. Thus the current registers must be saved as the state of the current process.

Process States: new -> ready -> CPU -> wait  

* new state : waiting for necessary resources (eg. CPU)
* read state : all resources except CPU. It is ready to run, just needs the CPU.
* CPU state : process is currently being executed by the CPU

A process starts off in the new state. When it is ready, it goes into the ready state. Notably, there many processes in the new state. When the CPU is available, a single process is chosen from all the processes that are in the ready state. How do you pick which process to run? There are numerous CPU scheduling algorithms to use to determine which ready process to move to the CPU state. One of the algorithms is a simple first in first out (FIFO) algorithm. There is also a wait state that the process can be in.

quantum: the slice of time given to a process by a CPU. This is nothing to do with physics


### Parent/Child Processes
Fork - parent process runs a child process  
While the child process is running, the parent process is in a suspended state. The child process can tell the parent process that it is done by called its exit() function.

### Process Control Block (PCB)
PCB is a general term for all computing. Specifically in UNIX, a Process Descriptor (PD) is a PCB. A process control block contains
* process ID
* register contents
* PC
* open files
* child processes
### CPU Scheduling
which processes gets the CPU
### Process Synchronization
includes race conditions, ...
### Process deadlock
An example of a process deadlock is:
* Process A is using file 1. It will release file 1 once it gets file 2.
* Process B is using file 2. It will release file 2 once it gets file 1.
In this scenario, Process A cannot move forward because it is waiting for resources that Process B has. It also won't release it's resource until it gets the resources B has. However, Process B cannot move forward either because it is waiting for resources that Process A has and won't release those resources until it gets the resources Process A has. Therefore, both cannot do anything.   

### Process Communication
Processes want to talk to each other.

They follow the producer/consumer structure often. This works by having one process (producer) load data into a memory buffer and then another process (consumer) reads from the same memory buffer.   

There are 3 main ways to do process communication
1. signals - send signals directly from a process to another
2. shared memory - set up variable that is shared. Thus buffer is in memory. This is what the example above describes. Note that the issue with this communication is that it requires Synchronization.
3. message passing - producer acts as memory. The cool thing about this communication technique is that it does not require synchronization.

### Memory Management
In this unit, we will learn
1. Memory space allocation and deallocation
2. trace of free and allocated memory space
  * there are multiple ways of doing. One way is with a linked list pointing to free memory
3. Virtual memory management
  * if you have 32 bit addresses, then you have 2^32 bytes in a given VM.
  * if you have 64 bit addresses, then you have 2^64 bytes in a given VM.
  * *memory size = page size*
  * paging, segmentation, or both can be used to get more memory.
### File Management
Before you can do proper file management, you have to design the disk layout. File management works closely with disk layout.
Some of the things we will talk about are:
1. disk layout
2. mass storage management
3. storage allocation/deallocation
  * contiguous or linked are 2 ways to organize blocks of storage
4. disk scheduling
5. free storage space management
### Protection & Security
encryption/decryption. RSA is the key one we will study

### System calls
There are 6 different system calls to note:
1. fork
2. execve
3. wait
4. exit
5. open
6. write
7. read - one of the params is fd
8. pipe - used for processes to communicate with each other.

API - set of operations to access system resources

### Batch Systems
Batch systems are those that run a queue of large jobs. Usually you submit a job to run and when it can, your appropriate job will run.

Batch systems are efficient for working with large datasets, however, there is a lot of overhead that comes with them. In fact, it is usually the case that the job itself is very quick, the reason the batch system is taking a long time to give you an answer is all the overhead tasks that it performs.

turnaround time - time from when you submit program to when you get output. For batch programs, your submitted program may be fast, but the turnaround time will slow because of all the overhead.

Batch systems are not good at developing programs, but efficient in running large long running processes.

multiprogramming - having multiple programs in memory at once. This allows for overlapping between CPU and IO

response time - time for starting program to *first* output. This is related to how much attention the program gets from the CPU.

timesharing - CPU can be switched rapidly from process to process. This doesn't reduce turnaround time, but reduce response time.

### Real-time Systems
For example: microprocessor in a DVD

Soft system: problems are tolerable (i.e. missing a few frames in a movie is tolerable, but such errors in another program (i.e. nuclear power plant) could be bad


### hand-held systems
For example: smartphones, ipads

usually run full-fledged operating systems, but have constrained screens, battery, memory. typically don’t use virtual memory management, mostly to save power (VM takes electrical power)
