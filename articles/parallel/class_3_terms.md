Class 3 Terms - High Throughput Terms
=======================================


### High throughput
using many resources over long periods of time to accomplish a computational task.

High-throughput computing requires large amounts of computing for long periods of time
High Performance computing requires large amounts of computing for short period of time

High throughput computing is about running long-running jobs that require lots of computing power. It is related to things like batch processing

### parameter space sweep
Given some problem with varying parameters, you go through all potential variations of the parameter in order to determine what the optimal output is.

For example, given f(x,y,z,a,b,c), you make a graph of f(...) given all combinations of x,y,z,a,b,c in order to determine what the optimal parameter values should be for your desired f value.

In computing, this literally can just mean running a program multiple times with varying parameters to get a feel for the output.

### embarrassingly parallel
little or no effort is needed to separate the problem into a number of parallel tasks. This occurs because there is little or no dependency between the parallel tasks so little communication is required between them.

An example is serving a static file to a user on a webserver. This can be done in parallel because you can initiate sending the user the file without having to worry about another user coming to the website and interacting with the first user.
