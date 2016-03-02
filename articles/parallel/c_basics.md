C Basics
============

### Pointers
Create a pointer

    int * ptr;
    char * mystr = malloc(sizeof(char) * len_of_str);

Make a pointer point to something

    *ptr = &something_on_heap; //* & is used to get pointer to something. is used to access value that pointer is pointing to.
    ptr = &something_on_heap // make ptr equal to address of something. Probably do not want to do this.

Dereference pointer

    int a = *ptr_to_num; //* deferences the ptr_to_num

Follow pointer of object

    struct Node{
      Node * left;
      int data;
    }
    //assume Node is properly created
    my_node.left = malloc(sizeof(Node)) // my node's left pointer now points to new node
    my_node->left = malloc(sizeof(Node)) // this is wrong. This would follow the left pointer. And then at that spot it would put in a new pointer instead of a Node object.

    (my_node->left).data = 3; //the -> follows the left pointer. This would allow you to left.data to be 3

    my_node.left.data; //ERROR. my_node.left is a pointer. pointers do not have data fields.    

#### Pointer arithmetic
C knows if something is a pointer so when you index or do any arithmetic involving that pointer, C uses the address that pointer is pointing to for the pointer and replaces all constants with sizeof(pointer)

    int * ptr = malloc(sizeof(int)); // assume malloc allocated memory location 0x01
    int b = ptr + 4; //ptr+4 = 0x01 + (0x04 * sizeof(int)) = 0x01 + (0x04 * 4) = 0x01 +  0x10 = 0x11 = 17

    int_ptr++ is the same as int_ptr + 1 so it would add 4 to the memory address pointed to by int_ptr

#### common type sizes
* char: 1 byte
* bool: 1 byte
* short: 2 bytes
* int: 4 bytes
* float: 4 bytes
* long: 4 bytes
* void * : 4 bytes
* T * : 4 bytes
* double: 8 bytes


???
int a = 3 + int_ptr + char_ptr; //???? is 3 multiplied by 4 or 1?



malloc() and the calloc() ??

#### cool idiom
subtract pointers to determine number of elements between those two in an array

#### null pointer versus void pointer
null pointer is a pointer that has value \x00.

void pointer is a pointer type that can point to anything

#### typed pointers
typed pointers refer to a specific type
void * : 4 bytes
T * : 4 bytes

### Heap management
(void * ) malloc(size_of_thing_you_want_to_create)
thus, you have to cast the pointer to a specific type!

int * p = (int * )malloc(sizeof(int));
??calloc

free(p); to free the space allocated by a pointer


There are two types of heaps - *fixed size blocks* and *buddy blocks*

##### fixed size blocks (also called memory pool allocation)
fixed-size blocks uses a free list of fixed-size blocks of memory all of the same size. If you want to allocate a large object, then you have to determine what x*block_size of memory is big enough for your object to fit in.

The problem with this type of allocation is that it leads to fragmented memory. This fragmentation occurs because of
* you may end up allocating tons of blocks that are too big for the small amount of memory that you actually need

##### buddy blocks
Memory is allocated in blocks of a certain factor of 2. This is good because you are differentiating between the objects that you need allocated and only allocating the appropriate amount for each block.

#### Auto variables
Auto variables are local variables that are allocated and deallocated when the variables goes out of scope. Autovariables are all stored on the stack.

#### parameter passing
parameters are passed in reverse order on the stack

#### stack
int main {
    foo(param1, param2)    
}

the stack for the above would look like:

1. param2
2. param1
3. ret_main
4. ebp_foo
5. foo_local_vars



##### Calling Convention
1. save caller saved registers (eax, ecx, edx)
2. push params
3. call some_func
  * a call consists of a push ret_addr and jmp some_func
  * ret_addr is just the next instruction after this call instruction
SUBROUTINE STARTS  
*follow callee convention*  
SUBROUTINE ENDS  
4. remove params from stack
5. return value of function is on eax
6. restore caller-saved registers(eax, ecx, edx)

#### Callee convention
1. push ebp; mov ebp, esp
2. allocate space for local variables
3. push calle-saved register (ebx, edi, esi)
*function body*  
4. put ret value at eax
5. pop callee save registers (ebx, edi, esi)
6. deallocate local variables
7. pop ebp
8. ret
  * ret consists of pop ret addr and jmp to it
  * ret_addr should already exist because of call in 3rd step of caller calling convention

### IO basics
To read stdin
    * scanf("format string", &var...)

to write to stdout
    * put(str)
    * printf(format_str, vars...)
#### open
FILE * fopen( const char * filename, const char * mode );
#### read
char * fgets( char * buf, int n, FILE * fp );
#### write
int fputs( const char * s, FILE * fp );
#### seek
int fseek ( FILE * stream, long int offset, int origin );
#### close
int fclose( FILE * fp );
#### sscanf
reads formatted input from a string

int sscanf(const char * str, const char * format, ...)

#### fprintf
sends formatted output to a stream.

int fprintf(FILE * stream, const char * format, ...)
