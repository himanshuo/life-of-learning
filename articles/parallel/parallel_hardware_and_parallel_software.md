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
