Lecture 3 - Computing Environments & System Resources
==========

2 types of computing environments
* traditional
* Networked

### Traditional computing environment
can have single or multiple processors (CPU)

### Networked computing environment
This is a broad category of computing environments that includes
* mobile computing
* distributed
  * examples: , NOS (network operating system)
    * DTOS (distributed operating system)
      * *high transparency* - you do not have to specify hostname when accessing files
    * NOS (network operating system)
      * this is not a true distributed system since it does required you to specify hostname
      * NOS is still distributed, but it just has a lower transparency
* client server paradigm
  * best example: http
* Peer to peer (p2p) (similar to distributed)
* cloud computing
  * Different types of cloud computing
    * SaaS (software as a service) - just have to give a service. Easiest form of cloud computing.
    * PaaS (platform as a service) - specify OS that you want to work on
    * IaaS (infrastructure as a service) - can install whatever OS you want to install based on your needs  

### 2 ways to access system resources
* command driven approach (interpreter based)
  * MS-DOS is one type, basically using command prompt
* Menu-driven (gui based)
  * most modern operating systems support this (uses mouse or touch screen)

### API
API (application programming interfaces) - include functions that are included to allow us to interact with a system

different systems have a different set of functions in their APIs.

Process/Thread management:
* CreateProcess(...) is the command for windows to make a new process
* fork() is the command for UNIX/linux, has no parameters
  * can be used to make a child process that is synchronized with the parent, wait functions are used for synchronization
* exit() - stop a process
* execup() is also designed for process management
* CreateThread(): windows for make thread
* pthread_create() for linux, p stands for POSIX package
* wait() is to wait for a process or thread to finish
* waitpid is to wait for a specific process or thread to finish

A sample usage is of these API functions is when you have a parent and a child process running. You want the parent to stop while the child is running.


In general, windows is more complicated

threads are not independent units, in that they operate within a single process and share resources such as memory space, creating possibilities such as producer/consumer
processes are independent

memory management:
* malloc() - allocate a portion of memory in the heap
* free() - free memory allocated

file management:
* open() - opens a file and returns file descriptor
* read() - reads info from a file
* write() - write info from a file

communications:
* socket programming
* STREAMS

##### Socket Programming
Socket: a data structure containing addresses like port number and IP address
  * IP address - the address of a host
  * port number - specific process on a the host

A couple of functions involving sockets:
* Socket() - creates an endpoint (socket) for communication. This returns a file descriptor for the socket.
* bind() - assigns a socket to an address. The socket is given an IP address and port number to potentially use.
* listen() - prepares for incoming connection. It basically just listens for communication.
* accept() - if communication came in, then a socket needs to establish a proper connection with it in order to properly send/recieve data. I assume this function represents the handshake between the two sockets.
* send/recieve - send/recive data


##### STREAMS
communication between applications, can be extended to networking

created for system V (old OS)

### Brief discussion of how to design an operating system:
##### monolithic approach
no structured approach to OS design. Just shove everything into a few files. Very efficient. Hard to adapt to changes. not good for larger systems, but very efficient


good for MS-DOS (old microsoft operating system, stands for microsoft disk operating system)


##### Layered Approach (bottom up)
* start with hardware
* make a layer of functions to change hardware
* then a higher layer
* ...
* user interface

once that layer is FULLY tested, make a next layer that can call previous functions, and make successive layers

not particularly efficient, because many successive calls are made from external layer to inner layers and finally to hardware, but easy to control complexity

Another caveat is that a lower level cannot use functions from a lower level.

Another issue is that this is very difficult to design. You have to know exactly what functions each layer will need beforehand as a lower level cannot use higher level functions.

*THE operating system*
This is an example OS using this layered approach.  

Features of THE:
* batch
* multiprogramming - multiple programs can be loaded into memory and run
* not multiuser


* layer 5: contains UI, essentially user layer
* Layer 4: contains user programs
* Layer 3: device i/o management
* Layer 2: communication between the OS and the console
  * console is under control of the operator
  * operator - loads program from IO device using console
* Layer 1: Memory and drum management
  * drum - is the old analog of disk
    * no seek time. thus very fast, but very expensive.
* Layer 0 : processor allocation and multiprocessing


*Layered system security (multics)*
* HW is surrounded by rings
* processes are assigned a ring.
* processes assigned a lower ring have higher privileges and can access more devices

### Microkernel Approach to designing OS
Microkernel was the approach Windows used at first, but then they changed to OO

kernel contains minimum set of functions. It does not have to be changed. This is good because if kernel is large (like in Linux), then changing something in the kernel requires shutting down entire OS.  

Microkernel Features:
* can set device registers
  * this is for *memory-mapped IO* - the memory and registers of the I/O devices are mapped to address values. This allows the CPU to manipulate the device.
* can switch CPU between processes (context switch)
  * process is given a tiny quantum of time to execute and then is pre-empted. In order to do this, the microkernel needs to be able to switch between processes.
  * processes are thus sharing the CPU
* can manipulate MMU (memory management unit)
  * maps from virtual address to physical address
* can capture hardware interrupts
* can pass system calls
  * must use message passing for microkernel

Examples of systems that were built using Microkernel approach:
* MACH
* TRU64
* WindowsNT

### Object Oriented Approach
OO is the common way to design Operating Systems nowadays



OO approach features:
* allows encapsulation, polymorphism, information hiding, inheritance
  * process manager can be treated as an object


Windows now uses OO now
* dll (dynamic link library) - contains many executable runnable files
* only smart part of OS is put into memory. Only necessary parts are included into memory via dll files.
* allow for dynamic loading and dynamic linking


# Process Management
process = program in execution

When process is running, the things contained in the execution environment are
* registers
* process id
* priority of process
  * for scheduling purposes

### fork
fork is a function in System Call in POSIX used to create a new child process
  * parameterless
  * returns pid (process id) to caller
    * pid is a small integer (between 0 and 65536=2^16)
  * fork creates the new process and automatically assigns it a pid


The child process is put somewhere in memory, and is a perfect copy of the parent. Both the child and parent have exactly the same code.  

For fork() to properly work, it has to identify which is the parent and which is the child. To this end, fork() return the pid of the child to the parent process and returns 0 in the code for the child. Now you can simply have a "if pid == 0" statement to determine whether the code to run should be run in the parent or the child.

Both processes have code that roughly says “if pid==0 do X else Y”, but this will only evaluate to true for the child, because pid = 0 was returned to the child


Child will often execute execvp, which takes in arg0 and args

The else clause usually has a wait clause followed by status because the parent must wait until the child process is all done

*example runcom function*
runcommand: a function that takes a command, parses it, and executes it

    int runcommand(char * cmd){
      //create variables
      char * argv[max_args];
      pid_t child_pid, c_pid; //pid_t is a custom type. It is just an int
      int child_status;

      //parse the command and put it inside argv
      parsecmd(cmd,argv);

      // create a child process
      child_pid = fork();

      if (child_pid == 0){ // cur code is running in child

        //run the child process.
        //execvp will delete the contents of the child process memory and replace it with whatever is needed to properly run the child command.
        //code after this should not run if execvp works correctly
        execvp(argv[0], argv);

        prinf(“unknown command!\n”);
        // should never be executed, because execvp will end the process itself
        exit();
      } else { // cur code is running in parent
        //wait until the child process is finished
        //note that the status of the child is given to the wait function
        c_pid = wait(&child_status);

        // you can check if the pid returned by wait is the same as the pid of the child you created.
        if(c_pid == child_pid){
          //ok
        } else{
          //something went wrong
        }

        printf(“child process =%d\n”, c_pid);
        return child_status;
    }
