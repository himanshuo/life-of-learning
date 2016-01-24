Parallel Hardware and Parallel Software
=======================================

# von Neumann architecture
This 'classical' architecture contains
* main memory,
* central processing unit (CPU) or processor or core,
* an interconnection between the memory and CPU

Main memory consists of collection of locations. Each location can store data and instructions.

CPU is divided into control unit and arithmetic and logic unit (ALU). Control unit determines which instructions to execute while the ALU actually executes the instruction.

Data in CPU and information about state of an executing program are stored in very fast storage - registers.  
Control unit has a special register called the program counter which stores the address of the next instruction to be executed.  
The interconnection is usually a bus - collection of parallel wires and some hardware controlling access to the wires. This interconnect allows you to read/fetch from memory and write/store into memory. The separation of memory and CPU is very important as it is a bottleneck in increasing performance.

# Processes, multitasking, and threads
OS manages hardware and software resources on a computer. It determines which programs can run and when they can run. It also controls allocation of memory to running programs and access to peripheral devices.

When a user runs a program, OS creates a *process* - instance of a computer program that is being executed. A process contains:
* execuatable machine language program
* block of memory that includes
  * executable code
  * call stack
  * heap
* descriptors of resources that OS has allocated to the process. For example, file descriptors
* Security information. For example, info on which hardware and software resources the process can access.
* information about the state of the process - is the process ready to run or is it waiting on some resource/register/memory.

Most modern OS's are *multitasking* - OS provides support for simultaneous execution of multiple programs.

Multitasking is possible for a system with a single core because each process runs for a small interval of time (*time slice*). After running program has executed for time slice, OS can run a different program.

*Blocking* is when a multitasking OS has a process that needs to wait for a resource and so will stop executing. While this process is stopped, the OS will run other processes.

*Threading* is a way programmers can divide their programs into more independent tasks with the property that when one thread is blocked, another thread can run. It's much faster to switch between threads than it is to switch between processes. This is because threads are "lighter weight" than processes. A thread can use the same IO devices and memory as the process. Thus sharing allows switching between threads to be faster.

If a process is the "master" thread of execution and threads are started and stopped by the process, then starting each thread is to "*fork* off the process" and to terminate the thread is to "*join* the process"

## Instruction-level parallelism

Instruction-level parallelism  attempts to run multiple processor components or *functional units* simultaneously run instructions. There are two main ways to do Instruction-level parallelism - *pipelining* and *multiple issue*.

### Pipelining
Pipelining is when you have functional units split up into stages and running them like an assembly line. Overall, a single functional unit will take longer, however, when combined with multiple others, the speed increases. Why? Because multiple stages can be done at once thus an instruction can start running before the previous instruction finishes.  

### Multiple Issue
Multiple issue is when you start multiple instructions simultaneously. Multiple issue processors replicate functional units and try to execute them simultaneously.

For example:
    for(i=0;i<1000;i++) z[i] = x[i] + y[i]
With multiple issue, you can calculate z[0] (aka the sum of x[0] and y[0]) and z[1] ( x[1]+y[1]) at the same time. Thus you end up doing two entire lines of code in the same time it normally takes to do one line of code.

If the functional units are scheduled at compile time, then a multiple issue system is using *static* multiple issue. If the functional units are schedule are run time, then the multiple issue system is using *dynamic* multiple issue. *Superscaler* = a computer architecture where several instructions are loaded at once and, as far as possible, are executed simultaneously. Dynamic multiple issue is superscaler.

In order to use multiple issue, a processor must find instructions that can be executed simultaneously. A key way to do so is through *speculation*, where you guess what instruction to run and how and if things work then yay but if things break then you need to have a strategy to fix things.

## Hardware multithreading

Hardware multithreading is just normal threads in a program. The things that are created by a process and you can have multiple threads and they can be run concurrently.

The good thing about *thread-level parallelism* is that it allows for a *courser-grained* (more general) parallelism than instruction level parallelism.

A key requirement of hardware multithreading is that you must be able to switch between threads quickly. This is the reason threads are awesome - because they are easy to switch between. This is why we run a gazillion concurrent threads more often than a gazillion concurrent processes.

 
