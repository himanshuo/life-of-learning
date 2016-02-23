Lecture 6
=============
KEY website to understand functinos: http://beej.us/guide/bgipc/output/html/multipage/pipes.html

### Parse
DIAGRAM on how to parse


Note that the diagram is not complete
* missing how to go to process state from various other states
* | and > are not complete

To check for end of line, use \n


### How to handle two pipes

Command

    cat scores | grep uva | cut c4

Header files ommited

code

    char * cat_args[] = {"cat", "scores", NULL};
    char * grep_args[] = {"grep", "uva", NULL};
    char * cut_args[] = {"cut", "c4", NULL};
    int pipes[4]; // total of 4
    pipe(pipes); // creates a pipe
    pipe(pipes+2); //creates another pipe
    //pipes[0] //read first |
    //pipes[1] // write first |
    //pipes[2] //read second |
    //pipes[3] //write second |



    if(fork()==0){
      dup2(pipes[1], 1);  //ouput -> pipe's input 
      //we close unused pipes so that default device used is the one that we want
      close(pipes[0]);
      close(pipes[1]);
      close(pipes[2]);
      close(pipes[3]);
      execvp(cat_args[0],cat_args);

    } else {
      if(fork()==0){
        dup2(pipes[0], 0);
        dup2(pipes[3], 1);
        close(pipes[0]);
        close(pipes[1]);
        close(pipes[2]);
        close(pipes[3]);
        execvp(grep_args[0], grep_args);
      } else {
        if (fork()==0){
          dups2(pipes[2], 0);
          close(pipes[0]);
          close(pipes[1]);
          close(pipes[2]);
          close(pipes[3]);
          execvp(cut_args[0], cut_args);
        }
      }
      for(i=0;i<3;i++){
        wait(&status)
      }
    }



### Shared memory approach for communication
For networked environments, you cannot use shared memory approach. In this case, you have to use socket programming.

Socket is just a connection on a end system. Socket contains
* port number
* ip address of host

Steps to use Socker
1. create a socket (SOCKET)
2. bind an address to the socket (BIND)
3. listen for a to the socket address for incoming request (LISTEN)
4. accept incoming request (ACCEPT)
5. read (RECIEVE) / write (SEND) to communicate with requester
6. disconnect communication (CLOSE)

Server:  
socker -> bind -> listen -> accept ->  receive /send  -> close   
Client:                     
socket -------------------> connect -> send / receive -> close   

server accept and client connect is called connection point

Usually the close command is initiated by the client. The server usually does not shut down/close.
