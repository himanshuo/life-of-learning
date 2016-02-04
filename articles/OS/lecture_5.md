Lecture 5
=============

DIAGRAM about different states of processes.
Also know examples

if you have a lot of new processes being created, then they remain in the new state. The intermediate term


long-term scheduler
short-term scheduler



### Linux
States

1.

init is first process.


stack is used for function calls
heap is for dynamically allocated memory
data

### Process Context-Switch
When CPU becomes available, you

This is fundamental to *multiprogramming* - multiple programs in memory.
To do process-context-switch, you have to
1. You have to save PCB for preempt-to-be process (done by dispatcher)
2. Then, restore PCB for dispatch-to-be process (done cooperatively by dispatcher and scheduler). Scheduler provides process.

Context has:
1. unique ID of the process (PID)
2. pointer to parent process
3. pointer to child processes
4. priority of process (for cpu scheduling)
5. register save area - area used to save registers
6. whether the processor is running or not
7. list of open files (file descriptor of each open file)
8. current position of stack pointer
9. current position of heap pointer
10. house keeping information (eg. cpu computing time. things that can help you determine priority. If high cpu computing time, then you have  )


### passive termination
exit()
### active termination
done by parent. parent checks status. you can call wait

### network
interprocess communication (IPC) is implemented on multiple nodes on network

2 approaches to do IPC:
1. shared memory approach.
2. message passing

*Shared Memory Approach*  
For example, say you have a producer-consumer problem. A buffer pool (DIAGRAM) can be used to

struct{

} item;
const in BUFFER_SIZE = 10;
item buffer[BUFFER_SIZE];
item next_item_produced, next_item_consumed;

the producer process:  
while(true){
  while( (in+1) % BUFFER_SIZE == out); //just waits. a way to waste cpu time.  [(n-1)+1] % 10 = 0. UNDERSTAND how to check fullness.  
  buffer[in] = next_item_produced;
  in = (in+1) % BUFFER_SIZE;
}

the consumer process:  
while(true){
  while(in == out);
  next_item_consumed = buffer[out]; // buffer is full so wait
  out = (out+1) % BUFFER_SIZE;
  // consume the item
}

in is a shared memory location. You don't want to do this.
you cannot use a counter either. Same shared memory issues occur.

*message passing*  
producer process:  

    while(true){
        // produce a item in item_message
        send(consumer_id, item_message);
    }

consumer process:  

    while(true){
      receive(producer, item_message);
      // consume the item in item_message
    }
