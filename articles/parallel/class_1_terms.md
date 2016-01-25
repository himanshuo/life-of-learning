Class 1 Terms
=================

## speedup
Speed up is a metric for improvement in performance between two systems processing the same problem. There are two different ways to measure speedup - *latency* and *throughput*.

Latency of an architecture is the reciprocal of the execution speed of a task. You can think of it as "time it takes to do 1 unit of work."
    L = 1/v = T/W
    where
    v = execution speed of task
    T = execution time of task
    W = execution workload of task
Latency is measure in seconds per unit of execution workload.

Throughput of an architecture is execution rate of a task. You can think of it as "work done per unit of time"."
    Q = pvA = pAW/T = pA/L
    where
    p = execution density (number of stages in an instruction pipeline)
    A = execution capacity (number of processors)
Throughput is measured in units of execution workload per second. Other frequent measures are intruction per cycle (IPC) and cycles per instruction (CPI).

Speedup has its own formula for latency and throughput.
    Speedup_Latency = old_latency / new_latency
    Speedup_Throughput = new_throughput / old_throughput

Note that speedup is supposed to be > 1 if things are going correctly.

## scaled speedup
## scalability
## top 500
## capability versus capacity computing
## amdahl's law
