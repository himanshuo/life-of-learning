Process Synchronization - ch 5
=======================

# Background

Processes can execute concurrently, but may be interrupted at any time therefore only partially completing execution.
Concurrent access to shared data may result in data inconsistency.

Maintaining data consistency requires ways to ensure orderly execution of cooperating processes.

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

      flag[i] = false;// process i is no longer ready to run the critical section
      remainder section  
    }while(true);

This meets the three required criteria to be a true solution to the critical-section problem.

* mutual exclusion - only 1 process enters into critical section at a time because Process i only enters critical section if either flag[j]=false (other processes are not in critical section) or if turn=i (it is process i's turn to enter critical section)
* progress - always looking for next process to come into critical section
* bounded-waiting - there is a list that determines how many processes can be waiting for chance to go to critical section

# Synchronization Hardware
many system provide hardware support for implementing the critical section code

## locks
All of these solutions are based on idea of **locking** - protecting critical regions via locks

But how do you have locks in a uniprocessor? you disable interrupts.
* currently current code would run without preemption

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
1. lock=false means that the lock is open and any process, including the current process, can use it.
2. the while(...) will return false and thus the critical section in the code can be run
3. the test_and_set function will make the value of lock be true thus signaling to others that the lock is not available for them to use
4. when the critical section is done, the lock will be set to false again signaling that the lock is available for use to others

When lock starts out as true
1. lock=true means that some other process is using the lock, thus the current process can't use it
2. the while(...) evaluates to true thus you cannot move forward
3. the * target= true line in the test_and_set code is basically doing nothing right now since the lock is already true.
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
Assume you are process i. Only n processes can be the requesting to go into critical section state. This state is called the 'waiting' state as defined by the the 'waiting' array in this program.  

    do{
      waiting[i] = true; //you are in waiting state
      // this section reduces to:
      // while(waiting[i] && test_and_set(&lock))
      // which basically blocks until both
      // 1. the lock is available for your use
      // 2. you are in the waiting state
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




# Mutex Locks
# Semaphores
# Classic problems of syncrhonization
# monitors
# Synchronization examples
# Alternative Approaches

# Topics Covered in Class
I covered the following topics in chapter 5. Race condition and atomic execution, hardware solution (enable and disable timer run out interrupt and test-and-set), software solution including Peterson's solution, semaphores (general/counting semaphores and binary semaphores), producer-consumer problem, readers-writers problem, dining philosophers problem, monitors (entry points, condition variables, and signaler queue), C code for the producer-consumer problem using condition variables.
