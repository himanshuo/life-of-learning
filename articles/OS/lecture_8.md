Lecture 8
=============

two privileged instructions

disable()
enable()


P1:
  while(1){
    disable();
    x++;
    enable();
  }

P2:
  while(1){
    disable();
    x--;
    enable();
  }


2 major problems
* malicious user can have control of cpu for ever
* solution doesnt work on multiprocessors 
