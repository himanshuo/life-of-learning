What is inside a Computer System?  
  * hardware
  * software
  * data
  * users

Layers of a computer system  
  * Applications (users and utilities)
  * Libraries
  * OS kernel
  * hardware

1) Applications use libraries  
1') Libraries can access OS kernel  
2) Applications can accesses OS kernel directly  
3) applications can access hardware


unix/linux: 1, 1'  
windows: 2  -   this makes it so that if you change kernel, then application must be changed as well. This is why windows is bad.


mode bit:    
determines whether the current instruction is for user or for kernel if user mode, then we check whether current user can do whatever action it wants to do


#### read(file1, block, memorybuf)
This simple seeming read function actually does a lot of stuff. The layers of abstractions are:
* compute position of block
* move block_size head to disk track ???couldnt read this???
* check for seek errors
* 4
* 5
* 6

#### virtualization
cpu creates 2 virtual cpu's
1 singe

timesharing:
spooling (for printing):


concurrancy
