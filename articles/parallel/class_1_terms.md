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

Throughput of an architecture is execution rate of a task. You can think of it as "work done per unit of time."
    Q = pvA = pAW/T = pA/L
    where
    p = execution density (number of stages in an instruction pipeline)
    A = execution capacity (number of processors)
Throughput is measured in units of execution workload per second. Other frequent measures are instruction per cycle (IPC) and cycles per instruction (CPI).

Speedup has its own formula for latency and throughput.
    Speedup_Latency = old_latency / new_latency
    Speedup_Throughput = new_throughput / old_throughput

Note that speedup is supposed to be > 1 if things are going correctly.

## scaled speedup
Conventional Speedup (from above) is basically just the reduction in execution time. If you run a slow version and fast version of a solution on a slow computer then you get a much greater speedup than if you made the comparison on a fast computer. Thus it makes it seem like your speedup is not as significant when you run it on a fast computer - normal speedup penalizes faster absolute speed. This is bad.

Scaled Speedup to the rescue! What scaled speedup does is simply scale the speedup equation so that you also consider the speed of the computer when determining speedup.

## scalability
Scalability is the capability of a system to handle growing amounts of work.

*Horizontal scaling* - add more nodes to a system
*Vertical scaling* - add more capability to a single node in the system

## top 500
list of top 500 super computers in the world

## capability versus capacity computing
Super computers generally aim for maximum in capability computing rather than capacity computing.

Capability computing is using the maximum computing power to solve a single large problem in the shortest amount of time.

Capacity computing is using efficient cost-effective computing power to solve a small number of somewhat large problems or a large number of small problems.

## amdahl's law
Amdahl's Law is a way to measure the maximum potential speedup of a system in terms of latency of a task with *fixed workload* of a system whose resources are improved.

An example - 30% of a program is sped up by half. Then the maximum speedup is only:
    old: (3/10)+(7/10) = 1
    new: (3/10) * (1/2) +(7/10) = 17/20
    speedup = 1 / 17/20 =  20/17 = 1.18
