Process Synchronization - ch 5
=======================

# Background

Processes can execute concurrently, but may be interrupted at any time therefore only partially completing execution.

Concurrent access to shared data may result in data inconsistency. Maintaining data consistency requires ways to ensure cooperating processes are executed in the order that they were intended.

For example, consider the producer/consumer problem using a counter:

    //producer
    while(true){
      while(counter == BUFFER_SIZE);
      buffer[in] = next_produced;
      in = (in+1) % BUFFER_SIZE;
      counter++;
    }

    //consumer
    while (true) {
	     while(counter == 0); // waste time
	     next_consumed = buffer[out];
	     out = (out + 1) % BUFFER_SIZE; 	
       counter--;
       // consume 'next_consumed'
    }

If you were to run this and the assembly instructions for counter++ and counter-- were interweaved, then you would race conditions.


# Critical-Section Problem
The critical-section problem is a hypothetical problem which models common synchronization problems.

The problem is that:
* consider n processes {p0, p1, p2, ...p_{n-1}}
* each process has a critical-section segment of code.
  * the critical-section may be changing common variables, updating table...
  * when a given process is in the critical-section, then no other may be in the critical section

Code illustrating problem

    do{
      entry section

      critical section

      exit section
      remainder section
    } while(true);


The solution to the critical section problem involves *mutual exclusion*, *progress*, and *bounded waiting*.


##### mutual exclusion
only 1 process can execute in the critical section at a time
##### progress
if no process is executing in the critical section and some processes want to enter the critical section, then the selection of which process to enter into the critical section *cannot be postponed indefinitely*.
##### Bounded Waiting
1. a process makes a request to enter the critical section
2. only x other processes are allowed to request to enter their critical sections
3. the process's request is granted

Bounded waiting is step 2 in this system.

Thus bounded waiting just means a limit on the number of processes that can be in the critical-section request state at a given time.

##### preemptive vs. non-preemptive
Note that the solution has two different approaches depending on whether the kernel is preemptive or non-preemptive.

preemptive = process can be preempted (moved from runnng state to blocked state) by kernel

non-preemptive = process runs until it exits kernel mode, blocks or voluntarily yields CPU.

Both are similar. Both refer to moving from running to blocked state. Difference is that preemptive means that *something other than the process (namely the kernel)* is allowed to preempt the process. non-preemptive means that only the process can do it.


A non-preemptive kernel is essentially free from race conditions because only the kernel process running can change the data while the kernel process is running.

### Petersons solution
assume that load and store are machine-language instructions that are atomic (cannot be interrupted).

The two processes share two variables
* int turn;
* boolean flag[2]

turn determines whose turn it is to enter critical section

flag determines indicates if process i is ready to enter critical section. flag[i]=true means that Process 1 is ready to enter critical section.

##### Code for Petersons solution
Note that this code would be inside process i.

Also assume that there are only two processes i and j. This can be extended to more processes.

    //this code would be inside process i
    do{
      flag[i] = true; //process i ready to enter critical section
      turn=j;//currently running process j
      while(flag[j] && turn ==j); //waste time while the current process is j and the turn is for j
      //we are waiting for the turn to be i or for process j to no longer be ready to run

      critical section

      flag[i] = false;// process i is no longer ready to run the critical
      remainder section  
    }while(true);

This meets the three required criteria to be a true solution to the critical-section problem.

* mutual exclusion - only 1 process enters into critical section at a time because Process i only enters critical section if either flag[j]=false (other processes are not in critical section) or if turn=i (it is process i's turn to enter critical section)
* progress - always looking for next process to come into critical section
* bounded-waiting - there is a list that determines how many processes can be waiting for chance to go to critical section. That list is the flag array. At the same time, turn has a limited range of values.

# Synchronization Hardware
many system provide hardware support for implementing the critical section code

## locks
All of these solutions are based on idea of **locking** - protecting critical regions via locks

But how do you have locks in a uniprocessor? you disable interrupts.
  * currently running code would execute without preemption
  * this is inefficient on multiprocessor systems.

You generally dont want to have concurrent code in uniprocessors because the overhead of locking (by disabling interrupts) makes things too slow. Plus this is not scalable.

How do you have locks normally? atomic hardware instructions.

atomic = non interruptible

Two ways to implement atomic hardware instructions
* test_and_set - test memory word and then set the value. this makes sure that you are only setting values after a set of tests to make sure everything will work is done
* compare-and-swap (CAS assembly instruction) - compare content of a memory location to a given value and, only if they are the same, modify the contents of that memory location to a given new value. The value to compare to will likely be the old value/assumed value at that memory location. If the values are different then the write fails.


##### sample code using locks

    do {
      acquire lock
        critical section
      release lock
        remainder section
    } while (TRUE);


### test_and_set lock

    //usage
    do {
          while (test_and_set(&lock)); // acquire lock. while you do not have a lock, do nothing.

          //critical section

          lock = false;
          //remainder section
       } while (true);

    //define the test_and_set function
    boolean test_and_set(boolean * target){
      boolean rv = * target;
      * target = true;
      return rv;
    }

To understand this, think a single process is trying to access the lock to run its critical section. There are numerous other processes that exist out there. The lock variable will either be true or false when the code for this process starts running.

When lock starts out as false
  1. `lock=false` means that the lock is open and any process, including the current process, can use it.
  2. the `while(...)` will return false and thus the critical section in the code can be run
  3. the test_and_set function will make the value of lock be true thus signaling to others that the lock is not available for them to use
  4. when the critical section is done, the lock will be set to false again signaling that the lock is available for use to others

When lock starts out as true
1. lock=true means that some other process is using the lock, thus the current process can't use it
2. the `while(...)` evaluates to true thus you cannot move forward
3. the `* target=` true line in the test_and_set code is basically doing nothing right now since the lock is already true.
4. the process will continue to check repeatedly until the lock becomes false.
5. once the lock is false, the section above on when locks starts out as false applies.  

### compare_and_swap lock

    //compare_and_swap function that manipulates the lock
    int compare_and_swap(int * value, int expected, int new_value){
      int temp = * value;
      if(* value == expected){
        * value = new_value;
      }
      return temp;
    }


    //usage
    do{
      while(compare_and_swap(&lock, 0, 1)!=0);
      //critical section
      lock=0;
      //remainder section
    }while(true);

Think a single process among many others. lock initialized to 0. 2 cases, lock=0, lock!=0.

when lock=0
1. lock=0 means that lock is open. no one is using lock. this process can use lock.
2. while(...) evaluates to false thus we can go into critical section
3. the compare_and_swap function will set the value to be the new_value which is non-zero. Therefore, signaling to others that the lock is being used. Therefore, not allowing others to run their critical sections
4. when done with the critical section, set lock=0 thus freeing it for other processes to use

when lock!=0
1. lock!=0 means lock is reserved by some other process
2. while(...) evaluates to true thus cannot go into critical section
3. lock!=0 thus lock will not be updated.
4. as soon as lock=0 thus other process frees it up, then while(...) will evaluate to true and allow you to use the lock and run your code

### Difference between test_and_set and compare_and_swap
test_and_set is for simply signaling if the lock is being used or not. This is evident by the fact that it uses a boolean value for the lock.

compare_and_swap allows you to know which process is locking. The expected value acts like an id for each process. This is evident by the fact that the lock is a number and is compared with the expected value.


### Bounded-waiting mutual exclusion with test_and_set
Assume you are process i. Only n processes can be in the 'requesting to go into critical section' state. This state is called the 'waiting' state as defined by the the 'waiting' array in this program.  

    //assume inside process i
    do{
      waiting[i] = true; //you are in waiting state
      // this section reduces to:
      // while(waiting[i] && test_and_set(&lock))
      // which basically blocks unless
      // 1. you are ready to go into the critical-section OR
      // 2. it is your turn to go into the critical-section
      key=true;         
      while(waiting[i] && key){
        key = test_and_set(&lock)
      }

      //critical section

      //find some other process that is not in the waiting state.
      j = (i+1) % n;
      while((j!=i) && !waiting[j]){
        j = (j+1) %n;
      }


    }while(true);


* this is thus basically the same as Peterson's solution but with different variable names and using a specific types of locks


# Mutex Locks
mutex lock has two key functions
  * `acquire()` is used to acquire a lock
    * blocks until you acquire the lock
  * `release()` is used to release a lock
  * both acquire and release are atomic

mutex locks are commonly implemented using **busy-waiting**. Busy-waiting simply runs a while loop that runs for however it takes the condition to be true. An example of a lock that uses busy-waiting is a **spinlock**

code for acquire

    acquire() {
     while (!available); //busy wait
     available = false;
    }

code for release

    release() {
           available = true;
        }

code to use mutex lock

    do {
        acquire lock
           critical section
        release lock
          remainder section
     } while (true);

# Semaphores
Semaphore S is just an integer variable that can only be accessed via `wait` and `signal` functions
  * wait is also called P()
  * signal is also called V()


code for wait

    wait(S){
      while (s<=0); //busy wait
      s--;
    }


code for signal

    signal(S){
      s++;
    }

wait(S)
  * if S >=1 -> value--, caller proceeds
  * if S <1 -> value--, calling thread is blocked and placed on a queue

signal(S)
  * if S >=1 -> S++, caller continues
  * if S < 1 -> there are 2 cases:
    * if no thread is waiting -> S is set to 1
    * if there are threads waiting -> one thread is released

##### counting semaphore
integer value can range over an unrestricted domain

##### binary semaphore
integer value can only be 0 or 1
  * this is the same as a mutex lock

### semaphore implementation
You can have the problem where the wait and signal functions are called at the same time.
  * this is basically just the critical-section problem
  * thus you can solve this problem using busy-waiting
  * however, busy-waiting is not optimal because it is only good for when you are spending a tiny amount of time in the critical section.
    * this is because busy-waiting uses up system resources instead of allowing other programs to run and use their resources.
  * solution is to use waiting queue

Each semaphore has an associated waiting queue. Each entry in the waiting queue has
  * value (int)
  * pointer to next record in list

Each semaphore also has two functions
  * block() - place process on appropriate waiting queue
  * wakeup() - remove one process from waiting queue and put it inside the ready queue


The idea is that there is a waiting queue for the semaphore. Each time you call wait(), block() may be called to put the calling thread inside the queue. Each time you call signal(), wakeup may be called to remove a process from the queue.

NOTE: the waiting queue is `different` from the list where processes are stored when the value inside wait() < 1. The waiting queue is specifically used to implement `blocking`.


code for a semaphore without busy wait

    typedef struct{
      int value;
      struct process * list;
    } semaphore

    wait(semaphore * S) {
      S->value--;
      if (S->value < 0) {
        add this process to S->list;
        block();
      }
    }

    signal(semaphore * S) {
       S->value++;
       if (S->value <= 0) {
          remove a process P from S->list;
          wakeup(P);
       }
    }

### Deadlock
Deadlock is when 2 or more processes are waiting indefinitely for an event that can be caused by only one of the waiting processes

Sample Deadlock code

    P0            P1
      wait(S)       wait(Q)
      wait(Q)       wait(S)
      ...           ...
      signal(S)     signal(Q)
      signal(Q)     signal(S)

* note that the two used semaphore Q and S in reverse order.
* one way to ensure no deadlock is to use semaphores in the same order in all processes/threads.

### Starvation
starvation is indefinite blocking

a process may never be removed from the semaphore queue in which it is suspended. This occurs because other processes keep coming in and getting the requested resource before this process does.

Notably, the resource is actually being used. If the other processes stop using the resource, the starved process will actually stop waiting and use the resource.

### priority inversion
priority inversion is when a higher priority task has a lock but then a lower priority process preempts the higher priority process. In doing so, the lower priority process has just taken over the higher priority process which is bad.

This is solved by the priority-inheritance protocol - increase priority of process A to max if any other process is waiting for a resource that process A has.
  * thus current process A will keep hold of the resource and all the other processes will remain in their appropriate priority


### Common Mistakes when using semaphores
Calling signal and then wait leads to problems

Calling wait and then wait again leads to problems

Forgetting to call wait and/or signal will lead to problems.


ALWAYS ALWAYS ALWAYS, call wait(S) and then signal(S)


### Semaphores can be implemented with Mutexes

variables used

    semaphore mutex; //initially=1
    semaphore next; //initially=0
    int next_count = 0;


wait() can be implemented using

    //body of wait()
    if(next_count>0)
      signal(next)
    else
      signal(mutex)

    x_count--;

# Classic problems of synchronization
### bounded-buffer problem
n buffers, each buffer can hold 1 item.

code sample

      // mutex = take control of buffer
      // full = how much you can add to buffer
      // empty = how many spots left to write in buffer
      semaphore mutex=1, full=0, empty=n;

      //producer
      do {
        wait(empty);  // write if there is room to write in buffer
        wait(mutex);  // take control of buffer
           ...
        // add next produced to the buffer
           ...
        signal(mutex); //release buffer
        signal(full);   // note that the buffer has values to read from
      } while (true);


      //consumer
      do{
        wait(full); //wait until you can consume from buffer
        signal(mutex); // take control of buffer
        //consume item from buffer
        signal(mutex); // release buffer
        signal(empty); // note that buffer has read a value thus an extra write spot has opened up
      }


### readers and writers problem
data set is shared among number of concurrent processes
  * readers - only read data
  * writers - read and write data

problem - allow multiple readers to read at the same time
  * only single writer can access shared data at the same time

The shared data among threads is going to be
  * data set
  * semaphore rw_mutex (initialized to 1)
    * means that current thread can either read or write to data set
  * semaphore mutex (initialized to 1)
    *
  * int read_count (initialized to 0)

writer code

    do {
          wait(rw_mutex);
               ...
          // writing is performed
               ...
          signal(rw_mutex);
     } while (true);


reader code

    do {
           wait(mutex); // mutex for read_count
           read_count++;
           if (read_count == 1) //if any reader, then can't write
              wait(rw_mutex);
           signal(mutex);
               ...
           // reading is performed
               ...
           wait(mutex);
           read count--;
           if (read_count == 0) //if no readers, then can write
           signal(rw_mutex);
           signal(mutex);
       } while (true);



### dining-philosophers problem
philosophers have 2 actions
  * eat
    * in order to eat, need to pick up chopsticks
  * think

assume 5 philosophers sitting in a circle

shared data
  * bowl of rice (data set)
  * chopsticks (semaphore S. initialized to 1. max of 5)

philosopher code

    //assume this code is run in philosopher i
    do {
      wait(chopstick[i]);     //pick up chopstick
      wait(chopstick[(i+1)%5])//pick up neighboring chopstick

      //eat

      signal(chopstick[i]);     //drop chopsticks
      signal(chopstick[(i+1)%5])

      //think
    }while(TRUE)

Scenario:

  1. think until the left fork is available; when it is, pick it up;
  2. think until the right fork is available; when it is, pick it up;
  3. when both forks are held, eat for a fixed amount of time;
  4. then, put the right fork down;
  5. then, put the left fork down;
  6. repeat from the beginning.

This is a normal, sensible scenario. However, think about it, step 1 checks out for all of them. None of them can get to step 2.

##### How to handle the deadlock for philosophers problem
Allow max 4 philosophers at a time

allow philosopher to pick up forks only if both left and right forks are available (picking up must be done in critical section)

Use an assymetric solution
  * odd numbered philosopher picks up first left chopstick and the right chopstick
  * even numbered philosopher picks up right and then left  

# monitors
monitors are a high level abstract data type used synchronization

![](process_synchronization-images/7be0421e3efd6e44ead2f6e5158ed7a0.png)
  * the oval is the monitor
  * only 1 process can be active within the monitor at one time
  * there are multiple queues that are input to the monitor
    * each condition variable defines a new queue input to the monitor
    * the condition variables are the shared variables

sample monitor code layout

    monitor monitor-name {
    	// shared variable declarations

      procedure P1 (…) { …. }
    	procedure Pn (…) {……}

        Initialization code (…) { … }
    }

### condition variables
condition variables are simply blocking points. They block until they are released.
  * wait() - when a process invokes wait(), the process is suspended until signal()
  * signal() - resumes one of the processes that invoked wait()
    * any one of the wait() processes can be invoked. There is no order.
    * if there is no wait() process, then signal() doesn't do anything



### (signal and wait) vs. (signal and continue)
if process Q calls wait() and then P later calls signal()
  * Q can start running now.
  * HOWEVER, P is already running. You can only have 1 thing running. So which one runs? P or Q? Both have equal logical reason to run.

Signal and wait - P waits for Q to be done

Signal and Continue - Q waits for P to be done

### Monitor Solution to Dining Philosophers

    monitor DiningPhilosophers{
      enum {THINKING, HUNGRY, EATING} state[5];
      condition self[5];

      initialization_code() {
        //everyone is thinking (not eating)
	       for (int i = 0; i < 5; i++)
	       state[i] = THINKING;
	     }

       void test (int i) {
         //if left and right is not eating and you are hungry
         // THEN signal that you can pickup
       	        if ((state[(i + 4) % 5] != EATING) &&
       	        (state[i] == HUNGRY) &&
       	        (state[(i + 1) % 5] != EATING) ) {
       	             state[i] = EATING ;
       		    self[i].signal () ;
       	        }
          }


      void pickup(int i){
        //say that you are hungry.
        // try to eat
        //if you are not eating, wait.
        state[i]=HUNGRY;
        test(i);
        if (state[i]!=EATING) self[i].wait;
      }

      void putdown(int i){
        state[i] = THINKING;
        //signal that the left and right have been put down
        test((i+4)%5);
        test((i+1)%5);
      }

    }


### conditional wait
x.wait(c) where x = condition, c= priority
  * process with lowest number (highest priority) is scheduled next



# Alternative Approaches
### transactional memory
sequence of read-write operation to memory that are atomic
  * a given function either works OR fails. No inbetween
    * if fails, then all of its changes are rolledback

### openmp
compiler directives and API that support parallel programming

    void update(int value)
       {
        #pragma omp critical
        {
          count += value
        }
        }

the pragma line is turned into code that openmp uses to do cool parallel stuff

### functional programming languages
functional languages do not maintain state

variables are all immutable and cannot change state once they have been assigned a value

# Topics Covered in Class
I covered the following topics in chapter 5. Race condition and atomic execution, hardware solution (enable and disable timer run out interrupt and test-and-set), software solution including Peterson's solution, semaphores (general/counting semaphores and binary semaphores), producer-consumer problem, readers-writers problem, dining philosophers problem, monitors (entry points, condition variables, and signaler queue), C code for the producer-consumer problem using condition variables.
