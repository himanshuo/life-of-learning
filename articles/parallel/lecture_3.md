Lecture 3
===================
### Optimization & timing
When giving timing information about a program, you always want to give metadata along with it. This metadata should include
* hardware
* compiler
* input data set
* cache
* test cases
* compiler flags
* motherboard
* load on machine
* time of day program is run

The load on a machine is the number of processes running on the machine. If you have 8 cores and 5 cores being used then a new process will automatically just use one of the three unused cores. However, if you have 8 cores and 9 processes already being run, then a new process will have to wait for something to finish before it can be run on a core.

##### Cool timing commands
* uptime
  * gives you load on a unix machine
* time <cmd>
  * gives you realtime (clock time), user time (time user spends on command), and sys time (time for given command only)
* gprof
    * profiler
    * gives you call graphs, time spent in each function, number of times each function/code is executed
    * example:
      * g++ -pg myprog.c -o myprog

### Optimizations
*Reasons to not optimize serial performance*
* time consuming  (humand time to rewrite program)
* makes code harder to read and debug
* can make it difficult to parallelize
* optimization on one architecture may not work on another

*when to optimize*
* code widely used in production
* projects with limited resources

*cycle to use when optimizing*
1. obtain code profile
2. tune next most time-sensitive section
3. if there is an increase in performance, go to step 1. Else, go to step 2.  

*2 pronged approach*
write optimizable code (code that can be run with -o, -o2, -o3 flags). Then, once code is debugged, try more aggressive optimizations.

*compiler will not do*
* reorder memory access
* factor/simplify large equations
* replace slow logic
* detect and remove excess work
* invert divisions

Since compiler wont do these things, you should do them to optimize yourself.

*How to help compiler with optimizations*
* always move linearly through memory
* factor/simplify equations
* remove excess work
* avoid branching inside loops

*optimization flags*
* o1: produce optimal code in shortest amount of time
* o2: more optimizations while not increasing program size
* o3: stronger optimizations and will increase program size if needed
* os: reduce program size

##### optimization tips
* avoid expensive math
* simplify math
* invest divisions where possible
* avoid branching in inner loops
* loop unrolling
* avoid excessive array lookups with temporary variables

*expensive math*
* addition, subtraction, multiplication are *very fast*
* division, exponentiation is *slow*
* square root is *very slow*

        a = sqrt(x^2 + y^2)
        if(a < b)
          ...
        SHOULD BE CHANGED TO
        if(x*x + y*y < b*b)
          ...

Doing this will remove the square root, exponentiation, and even creating a unnecessary temp variable.

*invert division*
    for(i:...)
      for(j:...)
        y = x / arr[i]
    SHOULD BE CHANGED TO
    for(i:...)
      temp = 1/arr[i]
      for(j:...)
        y = x * temp
This makes it so you are not multiplying at each loop instead of dividing. Of course, you still have to divide once per out for loop but it still is a lot faster to do this.

*branching inside loops and how to avoid it*

    for(...)
      if(...)
        ...
      Else
        ...
    SHOULD BE SUBSTITUTED FOR
    if(...)
      for(...)
        ...
    Else
      for(...)
        ...
Yes, this is more code and doesn't look as pretty, but because there is no internal branch, the compiler can turn the results of repeatedly used internal expressions into temporary variables. This is a key optimization.

*example equation simplification*

![](lecture_3/7efa8cfa9d2559d3d16544f0bd1cdb9e.png)

*inline keyword*
An inline function is a function whose body is placed inside its calling function rather than having to branch to it. 'inline' keyword gives a hint to the compiler to inline a function. However, this keyword might be obsolete nowadays since many compilers (when optimizing) already figure out and perform this optimization.

### cache effects
Memory is the bottleneck. So knowing and using the memory hierarchy properly is very important.

Two key things to know are that
* Data is fetched from memory as a cache line which is typically a 4,8,16,32 bit word.
* structs are stored together in memory

It is faster a lot of times to NOT make a struct and just put the desired featured of a grouping in an array because it leads to more cache hits. For example,

    GOAL: iterate through field a of all objects
    assume cache can only store 3 data items at a time
    using struct: [obj1_a obj1_b obj2_a obj2_b obj3_a obj3_b ]
    using array: [obj1_a obj2_a obj3_a obj1_b obj2_b obj3_b obj1_c obj2_c obj3_c]

the struct list would lead to multiple cache misses since the cache would only be able to store 1.5 objects at a time. However, in the array case, the cache will be able to store all field a values and thus be a lot faster.


### Parallel Hardware

*Interconnection Networks* connect processors to shared memory and to each other. There are two key types - *shared medium* and *switch medium*.

Shared medium is when you share resources. Some key features of it are:
* collisions require not sending the message
* only 1 message is sent at a time
* message is overbroadcasted
* each process listens to each message
* arbitration is decentralized

Switch medium is when you send multiple message at one time. Some key features of it are:
* point-to-point messages between pairs of processors


### Reasons to not optimize serial performance
* its time consuming
* not useful if code is not run often
* can make code harder to read and debug
* ooptimized code may not be able to scale

### When to optimize
* production code that is widely distributed and used
* often in the research community
* projects with limited resources
