# Virtual Memory

Virtual Memory makes it so that main memory functions as a cache for external storage. Virtual memory keeps only the active parts of many running programs. Inactive parts of running program are kept in a block of secondary storage called *swap space*.

Virtual Memory operates on blocks of data and instruction called *pages*. When a program is compiled, its pages are assigned *virtual page numbers*. Then, when the program is run, a *page table* is created that maps the virtual page numbers to physical addresses. Thus if the program requires a specific memory location, the resulting value that Virtual Memory gives it is from physical memory. The OS makes sure that memory used by one program does not interfere with memory used by another.
