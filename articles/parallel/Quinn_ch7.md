Quinn Chapter 7 - Performance Analysis
========================================
It is useful to be able to predict the performance of a parallel algorithm.
Some key things things to know:
* Amdahl's law  - determine whether a program merits parallelization.
* Gustafson-barsis's law  - way to evaluate the performance of a parallel algorithm.
* Karp-Flatt metric  - helps to decide whether the principle barrier to speedup is too much sequential code or parallel overhead.
* The isoefficiency metric  - way to evaluate scalability of a parallel algorithm executing on a parallel computer. It can help to choose the design that will achieve higher performance when the number of processors increases.

## Speedup and Efficiency

*Speedup = sequential execution time / parallel execution time*

The operations performed by a parallel algorithm can be put into three categories:
* Computations that must be done sequentially
* Computations that can be done in parallel
* Parallel overhead (communication operations and redundant computations)

Ψ(n,p) = speedup of some problem of size n on p processors   
σ(n) = sequential portion of the computation
φ(n) = parallel portion of computation
κ(n,p) = time required for parallel overhead

A sequential program operating on a *single* processor can only do one computation at a time: σ(n) + φ(n)  
This program does NOT have the κ(n,p) term because it does NOT require any interprocessor communications. In fact, it cannot do any interprocessor communications since it only has 1 processor. Also note that just because you have a single processor does not mean you cannot do any parallel computations - you can use threads.

The best possible parallel execution time occurs when all potential parallel portions of a program are made parallel. There still may be some sequential portions in the program, but it is impossible to make these parallel. These sequential portions take σ(n) time regardless of the number of processors added. 



## Amdahl's Law
### Limitations
### The Amdahl Effect
## Gustafson-barsis's Law
## Karp-Flatt Metric
## Isoefficiency Metric
##
