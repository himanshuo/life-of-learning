Lecture 14
==========
### project understanding

find max:  
15 8 15 34 91 24 17 68 (8 numbers)  
15   34    91    68    (4 numbers)  
34         91          (2 numbers)  
91  

log2(n) tells you how many rounds of comparisions.

should be inspecting 4096 numbers in inData.txt file.

ulimit - 5 512

upto 4096 threads


need to use barrier for synchronization. new barrier for each level of calculation.



sequence of functions to use:
* pthread_attr_init  
* pthread_create
* don't use pthread_join to block. Only use barrier.  
* sem_init
* sem_wait
* sem_post
* sem_destroy
* pthread_mutex_init
* pthread_mutex_wait
* pthread_mutex_unlock
* pthread_mutex_destroy
* pthread_cond_wait <- certainly need this.
* pthread_cond_signal
* pthread_cond_broadcast

he suggests to use mutex instead of semaphore.

must implement barrier, can't use pthread_barrier

can change stack size


test cases
* empty file
* small number of numbers
* all identical numbers
* all unique numbers
* all neg
* mix of neg/pos integers



### ???
1 MB = 2^20

1Kb = 2^10

2^20 / 2^10 = 2^10 = 1024 blocks

1024 bits
  0: free
  1: allocated

1024/8 = 128 characters

A = 3 KB
B = 2 KB
C = 5 KB
D = 1 KB
E = 5 KB


* memory release operation
  * bitwise AND opertion (&)
  * bit mask
* memory allocation
  * bitwise OR operation (|)
  * bit mask

in order to release block D: B[1]=B[1]&'11101111';

release block C: B[0] = B[0]&'11111000'; B[1]=B[1]&'00111111';


allocate first two KB of hole A: B[0]=B[0]|'11000000';


search k consecutive blocks
* logical right shift + bitwise AND operation
TEST=B[0]&'10000000';

right shift for the mask: 01000000
TEST=B[0]&'01000000';



### buddy system (static partitioning + dynamic partitioning)
16 blocks (0-15)

0,1 are buddies
2,3 are buddies

cannot combine 1 and 2 because they are not buddies
can combine 0 and 1 because they are buddies

0,1 & 2,3
4,5 & 6,7
12,13 & 14,15

0,1,2,3 & 4,5,6,7
8,9,10,11 & 12,13,14,15

0...15


IMAGE 
