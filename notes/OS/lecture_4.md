Lecture 4
=========

Review:
* a process creates a child process using fork()
* the child process is put somewhere in memory. It is an exact clone of the parent.
* you want both sets of code to have a "if pid == 0" line to determine whether the current code should be run for the parent or for the child.
* for the child, you generally want to use execvp to run some command. execvp will also delete this entire child process from memory and replace it with whatever command you are running with execvp.
* for the parent, you generally want to wait for the child process to finish. Once the child finished, you generally return the child processes status.

*Understanding Pipes*  
A pipe has  
* input read -
* input write
* output read
* output write

Default input devices
* 0 is for input
* 1 is for output
* 2 is for stderr


##### dup2 command
dup2 is a system call that duplicates one file descriptor, making them aliases, and then deleting the old file descriptor.
This becomes very useful when attempting to redirect output, as it automatically takes care of closing the old file descriptor, performing the redirection in one elegant command.

For example, if you wanted to redirect standard output to a file, then you would simply call dup2, providing the *open file descriptor* for the file as the first command and *1 (standard output)* as the second command.

    int dup2(int old_fd,int new_fd)


* int old_fd : source file descriptor. remains open after call to dup2
* int new_fd : destination file descriptor. This points to same file as old_fd
* return value : value of new_fd upon success

*sample code*

    #include <iostream>
    #include <unistd.h>
    #include <fcntl.h>

    using namespace std;

    int main()
    {
        //open a file
        int file = open("myfile.txt", O_APPEND | O_WRONLY);

        //Now we redirect standard output (1) to the file (file) using dup2
        dup2(file,1);


        //Now standard out has been redirected, we can write to
        // the file
        cout << "This will print in myfile.txt" << endl;

        return 0;
    }



Sample Input Command using Pipe:

    cat scores | grep uva

Sample Code to run input command:

    #include<stdio.h>
    #include<unistd.h> // constants & types
    #include<fcntl.h>  // file control
    #include<sts/types.h>
    #include<sys/stat.h>

    int main(int argc, char** argv) {
      //int pipefd[2] means that you have only 2 devices
      // pipefd[0] to read to the pipe (pipe's output read)
      // pipefd[1] to write to the pipe (pipe's input write)
      int pipefd[2];
      int pid;
      char* cat_arg[] = {“cat”, “scores”, null};
      char* grep_arg[] = {“grep”, “uva”, null};
      pipe(pipefd); //creates a pipe. will have to close input/output devices that you don't need.
      pid = fork(); // create a child process
      if(pid == 0) {  // child. child will run "grep uva"
        // pipefd[0]=pipe's output read, 0=standard input.
        // thus we are redirecting input (0) to the pipes output read (pipefd[0])
        // this means whenever there is standard input, it will NOT be handled like normal standard input, but rather be the output of the pipe.
        //grep reads the output of the pipe.
        dup2(pipefd[0], 0);   
        close(pipefd[1]);   //pipefd[1]=pipe's input write. We do not need to write to the pipe while running the grep statement.
        execvp(“grep”, grep_args); // execute
      } else { // parent. parent will run "cat scores"
        dup2(pipefd[1], 1); // standard output (1) will be written to the pipe's input. Thus the standard output can be used by some other command.  
        close(pipefd[0]);   // pipefd[0]=pipe's output read. We are not reading anything from the pipe.
        execvp(“cat”, cat_args);
      }
    }


I/O Redirection - not repeating header files

Command

    grep uva < scores > out

Code

    int inFile, outFile;
    char* grep_args[] = {“grep”, “uva”, null}
    inFile = open(“scores”, O_RDONLY);
    outFile = open(“out”, O_WRONLY);
    dup2(inFile, 0); // standard input (0) now is replaced by reading infile. Thus whenever you need to read standard input, you will instead read infile.
    dup2(outFile, 1); // standard output (1) now writes to outfile
    close(inFile);  // close the infile fd, because we already created a new file descriptor for reading the infile.
    close(outFile); // close the outfile fd, because we already created a new file descriptor for outputing
    execvp(“grep”, grep_args);

Command

    cat < redirect.c | wc > out

Code

    int main(int argc, char** argv) {
      int inFile, outFile;
      char* cat_args = {“cat”, NULL}
      char* wc_args = {“wc”, NULL}
      int status, i;
      pid_t first, second;
      int pipefd[2]; //pipefd[0]=input write, pipefd[1]=output read
      inFile = open(“redirect.c”, O_RDONLY);
      outFile = open(“out”, O_WRONLY);
      pipe(pipefd);
      first = fork();
      if(first == 0) {  
          dup2(inFile, 0); // standard input now reads from infile
          dup2(pipefd[1], 1); //  standard output now writes to pipe
          close(pipefd[0]); //do not need to read from pipe during cat command
          execvp(“cat”, cat_args);
      } else {
        second = fork();
        if(second == 0) {
          dup2(pipefd[0], 0); //standard input now reads from pipe
          dup2(outFile, 1); // standard output now goes outfile
          close(pipefd[1]); //dont need to write to pipe
          execvp(“wc”, wc_args);
        } else{
          // close all duplicated file descriptors
          close(inFile);
          close(outFile);
          close(pipefd[0]);
          close(pipefd[1]);
          for(int i = 0; i < 2; i++) wait(&status);
          print(“All done \n”);
        }
    }
  }

# Process Control Block (PCB)
Process control block stores all information necessary to the process. The process is in various states from start to finish and the PCB holds all the information needed for each state.


One of the things that PCB holds is the *state*. This allows the process to go through *state transitions*. 3 sample states:
* ready - process needs CPU time. The process at this point is already loaded
* running - process is using CPU already
* block - waiting for process to occur. It is waiting for something like a signal, event, ...

Sample state transitions:
1. new process has just been created.
  * When the PCB is created, you are in the new state.
  * all you have is the pid assigned to the process right now.
2. ready
  * ready is a state that means that the process is loaded but needs CPU time. However, it can be the case that there are too many processes that are ready. To handle this situation, you suspend a couple of them.
  * ready_a (active) is one subsection of ready state. It occurs when a process is ready and when you can immediately take this process and run it when CPU becomes available.
  * ready_s (suspended) is another subsection of ready state. It occurs when a process is ready but suspended.
  * you will have to switch between ready_a and ready_s depending on how many processes there are.
  * In addition, there are two key devices - CPU and IO. You can have two types of processes - CPU-bound and IO-bound. You want to balance CPU bound and IO bound processes.
    * CPU-bound - uses lots of CPU
    * IO-bound - process keeps doing lots of IO.
  * when you go from new to ready state, you get
    * priority queue
    * pid
    * resources
  * scheduler - makes schedule for ready processes to get CPU
  * dispatcher - performs context switch in order to allow a given process to properly run on CPU
    * context switch involves
      * restoring memory
      * restoring registers
      * restoring PCB
3. CPU state
  * when the process is running, it gets assigned a tiny amount of time called a *quantum* (usually 1/10th of a second)
  * note that *timesharing* is when multiple processes can run on the CPU by context-switching rapidly between them. The slice of time you get for your small timeshared portion is called a *quantum*
  * if your program is running and it runs out of your alloted quantum time then it is *preempted* from CPU state back to ready state. This is called *timer runout*.
    * there is an internal timer that keeps track of how long your process is running for in the CPU.
    * When you are supposed to be preempted, a *timer interrupt* is sent to the kernel which stops you from running further.  
    * the interrupt handler will move you from the CPU state to ready state.
    * in order for interrupt handler to do this, it has to perform a context-switch.
