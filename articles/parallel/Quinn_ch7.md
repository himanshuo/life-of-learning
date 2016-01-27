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



## Amdahl's Law
### Limitations
### The Amdahl Effect
## Gustafson-barsis's Law
## Karp-Flatt Metric
## Isoefficiency Metric
##
