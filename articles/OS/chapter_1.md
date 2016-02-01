Chapter 1 - Introduction
=========================
Operating system is a program that manages a computer's hardware.

![](chapter_1/4ec1bce8544f6e7ef5e0a7e9b99c7f58.png)

### What Operating Systems do

OS manages hardware, CPU, memory, IO device, and application programs. It provides an environment for all of them to work.

OS is the one program running at all times on the computer - usually called the *kernel*.  
*System programs* - programs associated with OS but necessarily part of the kernel.  
*application programs* - all programs not associated with OS.  
*middleware* - software frameworks that provide additional services to application developers.

### Computer-System Organization
When computer starts, a *bootstrap program* is run that starts the OS. The bootstrap program is stored in either read-only memory(ROM) or electrically erasable programmable read-only memory (EEPROM). EEPROM is also called *firmware*. This bootstrap program
* initializes CPU registers
* initializes devices controllers
* initializes memory
* load OS
* start executing OS
* starts system processes/system daemons
In order to accomplish all this, bootstrap program locates OS kernel and loads it into memory. Then, it executes the OS. It also does the same for system processes or system daemons. On UNIX, the first system process is "init" which starts many other daemons.

Once the OS and system daemons are loaded, the system waits for events. Events are signaled using *interrupts*. Hardware triggers interrupts by sending a signal to the CPU via the system bus. Software triggers interrupts using a *system call* operation. When an interrupt occurs, the OS stops its current work, and starts running the appropriate handler code for the interrupt. To determine the appropriate handler, the OS looks to a *interrupt vector* which maps interrupts to the address of the interrupt handler. The interrupt vector and the interrupt handlers are usually stored in the first hundred or so locations in memory.

### Storage Structure
CPU can load instructions only from memory so any programs that are run must be stored in memory. RAM is general purpose main memory. DRAM semiconductor technology is used for RAM. Static programs are stored in read-only memory(ROM) or electrically erasable programmable read-only memory (EEPROM).

2 key things to note about main memory:
* Main memory is usually too small to store all needed programs
* Main memory is volatile, meaning that it loses its contents when power is turned off

Secondary storage (including hard disk drives) handles both of the issues with main memory, however, it has much slower access time.  

![](chapter_1/1379b55c436db5037922adfb378ad380.png)


### IO  Structure
Most computer systems have multiple device controllers that are connected through a common bus. Each device controller is in charge of a specific type of device. A device controller has local buffer storage and a set of special-purpose regiters. The device controllers job is to move data between the device and the its local buffer storage. Normally, there is a device driver for each given device.

Steps taken for IO operation to occur:
1. device driver loads appropriate registers within device controller
2. device controller examines contents of these register to determine what action to take (eg. read a character from a keyboard)
3. controller starts the transfer of data from device to its local buffer
4. once data transfer is complete, device controller informs device driver via an interrupt that it has finished its operation.
5. device driver then returns control to the OS and usually returns the data or a pointer to the data.

For transferring large data, this byte by byte method is too slow. For transferring large data, there is direct memory access (DMA). DMA allows you to set up the buffers, pointer, and counters initially and then simply transfer entire blocks of data from the the buffer storage to the memory. Note that the device driver to device controller step is unchanged. This is only affect what happens when the data is already moved into the controllers local buffer. As each block is sent, an interrupt is sent to the OS. This is much faster than sending an interrupt per byte.

### Multiprocessor systems
Multiprocessor systems can do work in parallel. The key advantages to multiprocessor systems are
* *Increased Throughput* - Increasing the number of processors allow you to get more work done in less time. There is however some overhead time to make tasks parallel.
* *Economy of Scale* - multiprocessor systems cost less than equivalent multiple single-processor systems since they share storage, peripheral devices, and power supposed.
* *Increased Reliability* - failure of one processor will not halt the entire system, only slow it down.

Having multiple processors allows for graceful degradation - as hardware components fail, service is still provided.

There are two types of multiple-processor systems used today - *asymmetric multiprocessing* and *symmetric multiprocessing (SMP)*.

Asymmetric Multiprocessing is when each processor is assigned a specific task. A boss processor controls the system. This establishes a boss-worker relationship. The boss processor allocates work for the worker processors to do.

Symmetric Multiprocessing (SMP) is when all tasks are performed in parallel. This includes OS functions and user processes. SMP makes each processor a peer and the peers can communicate with each other. Each processor has its own registers and cache. They each share physical memory. The difficulty with this system is that IO must be carefully controlled to reach the correct processor. SMP is much more common than asymmetric multiprocessing.

Multiprocessing can cause systems to change from uniform memory access (UMA) to non-uniform memory access (NUMA). UMA is when any access to RAM from any CPU can take the same amount of time. NUMA is bad because it makes some parts of memory slower than others.

Multiple computing cores can be put on a single chip. This is called multicore. They can be more efficient than multiple chips with single cores because on-chip communication is faster than between-chip communication. In addition, one chip with multiple cores uses significantly less power than multiple single-core chips.  

##### Clustered Systems

Clustered systems gather multiple CPUs together. Clustered systems have two or more individual systems (called nodes) joined together. The nodes are loosely coupled. Each node may be a single processor system or a multicore system.

Clustering is used to provide high-availability service - service will continue even if some of the nodes fail. The idea is to add redundancy.

Clustering can be done symmetrically or asymmetrically. In symmetric clustering, one machine is in standby mode and simply monitors the other machine to see if it is functioning properly. If the other machine machines, the monitoring machine takes over the process. In asymmetric clustering, both machines run processes and monitor the other.

By parallelizing programs, you can run them in a large number of clusters and join their input together. This is embarrassingly parallel high performance computing.

A distributed lock manager (DLM) is included in some clustering technology in order to make sure conflicting instructions are not run and screw up the data.

### Operating System Structure
*Multiprogramming* is when you run multiple processes at the same time on a single processor. This is done by executing a part of a program, then part of another, and so on. To the user it appears that all programs are executing at the same time.

##### Key Terms
*firmware* = EEPROM memory from where startup programs and other system programs are stored.  
*middleware* = software frameworks that provide additional services to application developers.
*system processes/system daemons* = processes that are outside of the kernel that run the entire time the kernel is running.
*interrupts* = signal events in to the OS from either the hardware or the software  
*system call/monitor call* = software uses these operations to send interrupts to the OS.   
*interrupt vector* = maps interrupts to the address of the interrupt handler  
*RAM (random access memory)* = main memory  
*read-only memory(ROM)* = used for static programs such as bootstrap program   
*electrically erasable programmable read-only memory (EEPROM)* = used for static programs such as bootstrap program  
*Direct memory access* = method of transporting data from device controller to memory in block chunks in order to quickly transfer data from an IO device to the memory  
*graceful degradation* = ability to continue providing service proportional to the level of surviving hardware  
*fault tolerant* = systems that can suffer a failure of any single component and still continue operation  
*Multiprogramming* = organizing jobs (code and data)  so that the CPU always has one to execute
