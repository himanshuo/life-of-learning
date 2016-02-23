Lecture 4
==========

### Interconnection Networks
Interconnection networks connect processors to shared memory and processors to each other

Interconnection Networks pretty much drive how fast our programs will run

There are two interconnection media types
* shared medium
* switched medium

There are two opposite extremes in network design: *bus* and *fully connected network crossbar*. You want to find something in the middle to balance out both extremes

A bus is when you have a single connection between all processors. A bus provides
* good cost scalability - dont add any more connections with increasing number of processors
* poor performance - more processors end up having to share the same bus

A fully connected network (crossbar) is when you have a direct path between each component in the system. This yields
* good performance
* poor cost

A bus has bandwidth of 1, latency of 1.  
A crossbar has bandwidth of n, latency of 1.

bandwidth=rate of data transfer (bits/sec) from node to node

latency=time delay (sec/bit) between start to finish of data transfer

##### Shared medium
 ![](lecture_4/c42486e9b5b9dca05c1053192a5fb86b.png)

Key features of shared medium interconnection networks
* only one message is passed at a time
* messages are broadcast - each processor listens to every message
* arbitration (determining which value is correct when there is conflicting data) is decentralized
* collisions require resending messages


An example of shared medium interconnection network is ethernet.


##### Switched medium
![](lecture_4/2f0e2df5e936f7ecda501ce0969e7e29.png)

Key features of switched medium interconnection networks
* point-to-point messages between pairs of processors
* each processor has its own path to switch

Key advantages switched medium has over shared medium
* multiple messages can be sent simultaneously
* allows scaling of network to accommodate increase in processors

##### Switched Network Topologies
You can view a switched network as a graph where
* vertices's are processors or switches
* edges are communication paths

You end up with two kinds of topologies - *direct* and *indirect*.

Direct Topology is when the ratio of switch nodes to processor nodes is 1:1. This means that there is an equal number of switch nodes as processor nodes. For this to happen, you must have it be the case that every switch node is connected to either *1 processor node* and *atleast 1 other switch node*.

Indirect Topology is when the ratio of switch nodes to processor nodes is greater than 1:1. Thus you have more switch nodes than processor nodes. In this scenario, some switches just connect to other switches.

You can evaluate switch network topologies using the following characteristics
* diameter - largest distance between two nodes
* bisection width - minimum number of edges that must be removed to divide the network in half
* number of edges per node
* whether there is a constant edge length or not
* number of edges



### Evaluating Network Topologies
Again, the criteria is
* diameter - largest distance between two nodes
* bisection width - minimum number of edges that must be removed to divide the network in half
* number of edges per node
* whether there is a constant edge length or not
* number of edges

##### 2D meshes
![](lecture_4/86c0fb8ad4cf330b42322fbabe993afa.png)

Features:
* switches arranged into 2D lattice
* communication allowed only between neighboring switches
* direct topology
* the righthand side image is a variant where you have wraparound connections between switches on edge of mesh




### Architecture Types
* processor arrays - SIMD/ data parallel
* multiprocessors - shared memory
* multicomputers - distributed memory


most machines are hybrid

##### Processor arrays (SIMD/data parallel)
##### Multiprocessors (shared memory)
##### Multicomputers (distributed memory)
### Flynn's taxonomy


diameter - largest distance between two nodes
multiply this by latency to get measure
ex
--[]--[]--[]--[]-- n nodes. diameter is n/2, this is a ring
|____________ |

bisection width - if i cut the network in half, how many lines are in that half
- minimum number of edges needed to do this
- cross bar has best
number of edges per node
constant edge length or not
number of edges

Cost vs Performance
bus - good cost scalability, poor performance scalability
crossbar - good perfomance scalability, poor cost perfomance
Want to strike balance between these

2-D meshes
can be grid or taurus
taurus is rings to shorten diameter
for grid - diameter is dimension * width-1

Vector computers
instrcution set can operate on vectors as well as scalars
two ways to implement
pipelining (used for small scale)
processor arrays (used for large scale)
control units were expensive a long time ago. so just one control unit and lots of alus
performance (work per time unit)
depends on speed and if vector size fits
when you hit an if statement. always execute both
processor units set flag for which one
so we see speed drops for conditional code
control units are cheap now

Multiprocessors
multiple cpus with a single shared memory
same address on two different cpus refers to the same memory location
can be built from common cpus, naturally support many users
efficiency is maintained with conditional code

Centralized - scale out many cpus with a bus
straightforward, share same primary memory

private data - used only by single processor
shared data - used by multiple
challenges: synchronization, cache coherence.
cache coherence, need to replicate data across caches
so we have have a cache control protocol to make sure value is the same across by checking if it changes
if it changes, remove old values from other caches

Distributed
distribute memory space
each cpu has its own memory by partitioning memory for each
based on address to write, cpu can tell which cpu that memory should go to.
so it uses interconnection network
no global memory
effects:
increases aggregate memory bandwidth. scales linearly
some accesses are fast and local, other longer and subject to contention
lower average memory access time
allows greater number of processors
also called numas - non uniform memory access

cache coherence problem
cant snoop traffic over single bus
harder to do

multicomputer - distributed memory machine
same address on differenc processors referes to different physical memory locations
processors interact through message passing
disadvantages
single point of failure
primitive backends make debugging hard

Commodity cluster
very similar to multicomputer
low latency high bandwidth - mc
high latency low bandwidth - cluster

Flynn's taxomny
classify machine based on num instruction streams and num data streams. single or multiple
combinations:
SISD, SIMD, MISD, MIMD
