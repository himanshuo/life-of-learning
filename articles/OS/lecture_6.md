Lecture 6
=============
KEY website to understand functions: http://beej.us/guide/bgipc/output/html/multipage/pipes.html

### Parsing project 1
The diagram on how to parse the input for project 1 can be found online.

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
For networked environments, you cannot use shared memory approach because there is no shared memory between them - they are on completely different computers. In this case, you have to use socket programming.

Socket is just a connection on a end system. There is one on the source end and destination end of a connection. Socket contains
* ip address of host
* port number of host
  * port number is just the id of a process running on the host

Example: HTTP server
* HTTP default port number is 80
* When your browser is talking to a server to a given server, the socket it contacts has
  * ip address: ip address of host
  * port number: 80
* on the host, http process has id 80
* thus it makes sense that the port number you connect to is 80.

Maybe the above example is confusing. The key point I am trying to get across is: **the port number is the process id** on a given host.


Steps to use Socket
1. create a socket (SOCKET)
  * creates a socket data structure
2. bind an address to the socket (BIND)
  * put ip address and port number into the socket data structure
3. listen for a to the socket address for incoming request (LISTEN)
4. accept incoming request (ACCEPT)
5. read (RECIEVE) / write (SEND) to communicate with requester
6. disconnect communication (CLOSE)


Server flow versus client flow:    
![](lecture_6/2e16abfb68bb26da4c93ca57ff5bba68.png)

Key ideas for image
* The blue is the server flow.
* The green is the client flow.
* the server is setup separately and follows the start of its flow.
* client is set up separately and doesnt have a lot of steps to follow
* client initiates entire connection
* once the server and client make the connection and accept the connection, you can send/recieve data between the two.
* you can send/receive data from client to server and from server to client.


Server accept and Client connect point is called **connection point**

From clients perspective,
* send = write to server  
* receive = read from server

From servers perspective,
* send =  write to client
* receive = read from client


Usually the close command is initiated by the client. The server usually does not shut down/close.
