Creating a Virtual Memory Manager
==================================
A virtual memory manager maps a logical address to a physical address.

Segmentation allows the physical address space of a process to be non-contiguous. Paging is another memory-management scheme that does the same this. However, paging avoids external fragmentation and the need for compaction, whereas segmentation does not.

Paging solves the problem of fitting memory chunks of varying sizes onto the backing store.

### Basic Paging
Physical memory is broken into fixed-sized blocks called **frames**, and breaking logical memory into blocks of the *same size* called **pages**.

To load a process, load the process's pages onto memory from the source(backing store or file system).

The logical address space is totally separate from the physical address space. However, the backing store is divided into fized-size blocks that are the same size as the memory frames. 
