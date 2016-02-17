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


Sample Input Command using Pipe:

    cat scores | grep uva

Sample Code to run input command:

    #include<stdio.h>
    #include<unistd.h> // constants & types
    #include<fcntl.h>  // file control
    #include<sts/types.h>
    #include<sys/stat.h>

    int main(int argc, char** argv) {
      int pipefd[2];
      int pid;
      char* cat_arg[] = {“cat”, “scores”, null};
      char* grep_arg[] = {“grep”, “uva”, null};
      pipe(pipefd);
      pid = fork(); // create a child process
      if(pid == 0) {
        dup2(pipefd[0], 0); // oldfd, newfd
        close(pipefd[1]);
        execvp(“grep”, grep_args); // execute
      } else {
        dup2(pipefd[1], 1);
        close(pipefd[0]);
        execvp(“cat”, cat_args);
      }
    }


    I/O Redirection - not repeating header files
    grep uva < scores > out
    int inFile, outFile;
    char* grep_args[] = {“grep”, “uva”, null}
    inFile = open(“scores”, O_RDONLY);
    outFile = open(“out”, O_WRONLY);
    dup2(inFile, 0); => links scores to input device
    dup2(outFile, 1); => links out to the output device
    close(inFile);
    close(outFile);
    execvp(“grep”, grep_args);
    cat < redirect.c | wc > out
    int main(int argc, char** argv) {
    int inFile, outFile;
    char* cat_args = {“cat”, NULL}
    char* wc_args = {“wc”, NULL}
    int status, i;
    pid_t first, second;
    int pipefd[2];
    inFile = open(“redirect.c”, O_RDONLY);
    outFile = open(“out”, O_WRONLY);
    pipe(pipefd);
    first = fork();
    if(first == 0) {
         dup2(inFile, 0);
         dup2(pipefd[1], 1);
         close(pipefd[0]);
         execvp(“cat”, cat_args);
    }
    else if(first > 0) {  
    second = fork();
         if(second == 0) {
              dup2(pipefd[0], 0);
              dup2(outFile, 1);
              close(pipefd[1]);
              execvp(“wc”, wc_args);
         }
         else if(second > 0) {
              close(inFile);
              close(outFile);
              close(pipefd[0]);
              close(pipefd[1]);
              for(int i = 0; i < 2; i++)
                   wait(&status)
              print(“All done \n”);
         }
    }



Chapter 3
Process control block stores all information pertanent to the process
(1) state transition
there are 3 states, ready - needs CPU time it’s already loaded, running - when the process is using the CPU already, block - waiting for process to occur, waiting for something (signal, event, etc)
new indicates a new process has been created,
