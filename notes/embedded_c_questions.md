Top Questions for Interviews
http://a4academics.com/interview-questions/57-c-plus-plus/722-embedded-c-interview-questions?showall=&limitstart=
https://medium.com/@tattwei46/how-to-prepare-for-embedded-c-interview-9d499e9f9963
	1. What is the Volatile Keyword
		- Tells the compiler not to optimize variable usage
		- I.e. multiple reads from the same presumably unchanged memory location
		- Good for the use of memory mapped peripherals and multithreaded applications with shared locations
	2. Can variables be const and volatile?
		- Const ensures the value of variable declared cannot change internally
		- However it can be changed by external sources so volatile can be used as well
	3. Can a pointer be volatile
		- Typical declaration volatile int *p will say that the value at p will be volatile not p itself.
	4. What are the Size of various pointers and data types
		- Character is 1 byte
		- Integer is typically 4 bytes
		- Pointers are 8 bytes on 64 bit systems and 4 bytes of 32 bit systems
	5. Void pointer and its usage?
		- Variable of any type so it is a generic pointer
		- Used for struct pointers and standard data, even function pointers to an extent
		- Need to cast before using to access anything
	6. What is a NULL pointer and what is its use
		- Pointer to address 0x0 which is typically unused in user space
		- Means it does not point to any valid location
		- Pointers are set to null when we don't want it to be in charge of anything anymore.
		- We can use it to verify if a pointer is to a valid location or not
		- Its lazy design though and should be replaced with a better type as NULL can cause many issues in production. Empty and None are better replacements for the concepts of NULL.
	7. What is ISR
		- Interrupt Service Routine is an interrupt handler. A call back subroutine that is called when an interrupt is encountered.
		- Kernel code can trap into ISRs when desired to call certain behaviors
		- Upcalls can occur into kernel and user code when an ISR runs its routine when a hardware action has completed such as a disk read.
		- Therefore can allow for asynchronous execution
		- Vector table in the basic firmware is used to map Interrupt Requests by number/id to the correct ISR
	8. What is the return type of ISR
		- Has no return type as there is no real caller of the code itself.
		- Thus the concept of traps and upcalls are used because there is no code caller just hardware.
	9. What is interrupt latency?
		- Time required for an ISR to respond to an interrupt
	10. How to reduce interrupt latency?
		- Minimized by writing short ISR's and not delaying interrupts for more time
	11. What functions can be used inside ISRs?
		- Any function can be sued as long as that function is not invoked from other portions of the code
	12. Can printf be used in ISRs?
		- No. It is not reentrant, it is thread safe, and uses dynamic memory allocation which affects a lot of ISR speed.
	13. Can we breakpoint inside an ISR?
		- Not a good idea due to the speed hit. Logs are a better option.
	14. Can static variables be declared in a header file?
		- Static variables cannot be declared without definition.
		- Can be declared in header however a private copy will be made in each source file that includes it
	15. Are count up or Down loops better?
		- Count down loops are much better. The comparison to zero can be optimized out much easier as assembly have instructions for cmpz 
	16. What are inline functions?
		- ARM compilers support inline functions with __inline keyword
		- Small definitions and function body is substitued into each call to the function. 
		- Allows skipping of argument passing and stack maintenance
		- Increases code size but has faster execution.
	17. Can include files be nested?
		- Yes any number of times but they cannot be included twice!
		- May run out of stack room with increased number of headers however
	18. What are the uses of the static keyword?
		- Used with variables and functions
		- Static variable will be of static storage class and within a function it maintains its value between calls of the function
		- Variable declared as static within a file, scope will be in that file but cannot be accessed by other files.
	19. What are the uses of the keyword volatile?
		- Prevents compiler optimization which can change unexpectedly externally to compilers knowledge
		- Volatile variable will not be cached by the compiler in a register
	20. What is the top and bottom half of a kernel?
		- To handle an interrupt a lot of work needs to be done. This conflicts with need for speed of execution.
		- Linux splits handler into a top half and bottom half. 
		- Top half responds to the interrupt. Bottom half is a routine that is scheduled by top half to be executed at a safe time.
		- All interrupts are enabled during the execution of the bottom half. 
		- Top half saves device data to specific buffer schedules the bottom half and exits.
		- Top half can service new interrupt while the bottom half works on the previous interrupt.
	21. RISC vs CISC processors
		- RISC (Reduced Instruction Set Computer) has limited number of ASM instructions.
			§ Less hardware, fewer transistors and cheaper to manufacture.
		- CISC (Complex Instruction Set Computer) capable of executing multiple operations through single instructions compared to RISC
			§ More hardware, higher transistor count, more expensive
		- Most processors are RISC under the hood with CISC translation layer to use more instructions
	22. What is an RTOS?
		- Real time operating system. 
		- Scheduler in RTOS designed to provide a predictable execution pattern in a defined time period.
		- To keep real time requirements scheduler must be predictable
			§ Context switching latency needs to be short
			§ Interrupt latency needs to be short
			§ Interrupt dispatch latency need be short
			§ Reliable and time bounded inter process mechanisms
			§ Should support kernel preemption 
	23. Soft vs. Hard RTOS?
		- Hard needs to strictly make the deadline associated with a task; single failure to do so is considered complete system failure
		- Soft allows missing deadlines; real time tasks get priority over other tasks and retains it until completed.
	24. What type of Scheduling does an RTOS use?
		- Uses preemptive scheduling. Higher priority tasks can interrupt a running process to be resumed later.
	25. What is Priority Inversion?
		- When two tasks share a resource higher priority task will run first.
		- If lower priority task using shared resource when high priority task becomes ready the higher priority must wait for the lower to finish.
		- So because a higher priority task is waiting for a lower one this is inversion.
	26. What is priority Inheritance?
		-  Solution to inversion
		- Process waiting for any resource which has resource lock will have maximum priority
		- When one or more high priority jobs are blocked by a job the original priority is ignored and execution of the critical section will be assigned to job with highest elevated priority.
		- Return to normal priority when critical section is exited.
	27. How many types of IPC (Inter process communication) mechanism can you name?
		- Pipes
		- Named pipes (FIFO)
		- Semaphores
		- Shared memory
		- Message queue
		- Socket
	28. What is a spin lock?
		- When a resource is in use a thread that wants to use it can repeatedly check if the resource has become available.
		- Locks the thread in a spinning fashion where no useful work can be done.
	29. What is a semaphore?
		- Variable or data type which controls access to a common resource by multiple processes.
		- Two types:
			§ Binary semaphore (can only have 0 or 1 values) Set to 1 by process in charge when resource is available.
			§ Counting semaphore - value greater than one used to control access to a pool of resources.
	30. What are the main differences between Mutex and Semaphore?
		- Mutual exclusion and synchronization can be used by binary semaphore where mutex is used only for mutual exclusion
		- Mutex can be released by same thread that acquired it. Semaphore values can be changed by all threads.
		- ISR mutex cannot be used.
		- Semaphores can be used to sync two unrelated processes trying to access the same resource.
		- Semaphores can act as mutexes but the reverse is not true.
	31. What is virtual memory?
		- Program is able to pretend that it has all of the memory space.
		- Kernel / Hardware will actually take care of mapping it to physical memory using paging in most cases
		- Virtual memory can allow for hard drive swap space.
	32. What is kernel paging?
		- Memory management scheme that allows computers to store and receive data from the secondary memory storage when needed in primary storage
		- OS retrieves data from secondary storage in blocks called pages
		- Allows physical address space of a process to be non-continuous
	33. Can structures be passed to functions by value
		- Possible but not encouraged
		- Values changed in the structure will not be reflected in the caller
		- Large structures being copied will eat RAM on stack
	34. Why can arrays not be passed by value?
		- Arrays are not real in C. They are simply pointers so it is not possible.
	35. Macro and Inline function pros and cons
		- Overhead for the argument passing is reduced in in-lined.
		- Macro functions allow type insensitive functions
		- Macro's cannot do validation checks.
		- Both increase size of executable.
	36. What happens with recursive inline functions
		- Reduces the overhead of saving context on the stack.
		- Only a suggestion so the compiler does not have to guarantee this.
		- It may or may not inline it for several levels deep.
	37. Why do passing macro into a macro have difficulties?
		- Expands only one layer at a time and screws things up.
	38. ++*ip increments what?
		- Increments the value to which ip points to not the address
	39. Declare manifest const that returns number of seconds in a year using preprocessor?
		- #define SECONDS_IN_YEAR (60UL*60UL*24UL*365UL)
	40. Definitions and matching code
		- An integer
		- Pointer to integer
		- Pointer to pointer of integer
		- Array of ten integers
		- Array of ten pointers to integers
		- Pointer to array of ten integers
		- Pointer to function that takes in integer and returns integer
		- Pass array of ten pointers to function that takes an integer argument and returns an integer

		- Int a; 
		- Int *a; 
		- Int **a;
		- Int a[10];
		- Int *a[10];
		- Int (*a)[10];
		- Int (*a)(int);
		- Int (*a[10])(int);
	41. Why are typedef's preferred to defines for structs?
		- When you try to declare two struct pointers together the define will only work for the first declaration in the chain not the second
		- #define B struct A*
		- Typedef struct A *C;
		- C p1, p2; -> this works    B p1, p2; -> does not work
	42. What does malloc(0) return?
		- Returns a valid pointer but allocates no space for the pointer
	43. What is purpose of keyword const?
		- Read only value for the code
		- Value at the end of a const pointer can be changed
		- Const int a; or int const a;   -> the same
		- Const int *p    -> pointer to constant integer
		- Int * const p   -> address is constant
		- Int const * a const    -> address and value are const
	44. Blah
	45. How do you tell if a processor is big or little endian?
		- Can use a union or core dump file to test in code or by seeing the memory at a known location
	46. What is the concatination operator?
		- ## used to concatenate two arguments 
		- Literally in code the arguments are concatinated not the values at them
		- For example you can construct a variable name with the ## operator
		- Int var = 15;
		- Printf("&d", v##ar);
		- Prints var.
	47. How do you code an infinite loop?
		- While(1) {}
		- For (;;) {}
		- Do {} while(1);
		- Loop:
		- Goto Loop
	48. What does fork do?
		- Splits the current thread function from that point onward until the end of the function
	49. What is forward referencing with respect to pointers in C?
		- Used when a pointer is declared and the compiler reserves memory for the pointer, but the variable of the data type is not defined to which the pointer points to.
		- Ex:
			§ Struct A *p;
			§ Struct A { // members };
	50. How is list manipulation function written which accepts elements of any kind?
		- Uses void pointers when you don't know the stored data type.
	51. How can you define a structure with bit field members?
		- Use the ":" operator to define the width of bits of a char
		- Ex:
			§ Struct A {
				□ Char c1 : 3;
				□ Char c2 : 4;
				□ Char c3 : 1;
			§ }
	52. Function to take a byte and field in the byte return value of the field in that bit
		- Int GetFieldValue(int byte, int field) {
			§ Byte = byte >> field;
			§ Byte = byte &0x01;
			§ Return byte }
	53. What parameters decide the size of data type for a processor?
		- Compiler chooses but OS will limit by what is supported
	54. What is the job of the preprocessor, compiler, assembler and linker?
		1. Processes include files, conditional compilation instructions, and macros
		2. Compilation takes the output of the preprocessor and source code and generates the assembler source code.
		3. Assembly takes source code and provides an assembly listing with offsets. Output stored in object file.
		4. Listing/Linker takes output files and libraries as input combines to produce a single file. Resolves external symbols, assigns final addresses to all the functions / procedures, variables. Then revises code and data to reflect new addresses.
	55. What is the difference between static linking and dynamic linking
		- Static Linking: all library modules used in the program are placed in the final exe making it larger in size. Done by the linker. If modules are modified after linking we need to re-compile. Don't need to worry about incompatibility
		- Dynamic Linking: only the names of the module are used in the exe. Actual linking is done at run time when a program and libraries are both in memory. Exe's are smaller in size. Modification does not force re-compilation. Introduces possible compatibility issues.
			§ Allows DLLs so shared libraries.
	56. What is the purpose of #error
		- Used to throw errors in compilation
		- Sanity checks for files and using debug options. 
	57. Sometimes need to set value of 0xaa55 to 0x67a9 (magic value for bootloader)
		- Int *ptr;
		- Ptr = (int *)0x67a9;
		- *ptr = 0xaa55;
	58. What is the Watchdog timer
		- Time device with a predefined interval
		- During the interval something needs to happen or else the device generates a time out signal
		- Resetting the timer is called kicking the watch dog.
		- Resets to the original state when it times out.
	59. Why does ++n go faster than n+1?
		- ++ can be done with an ASM instruction such as INR
	60. What is a wild pointer?
		- Pointer than is not initialized to any valid address or NULL is considered wild.
		- Int *p;
		- *p = 20;
	61. What is a dangling pointer?
		- When a pointer that has been de-allocated or freed and the pointer is not set to NULL.
		- It can still contain the address of memory it should not be able to access. So can give errors upon dereferencing. 
	62. What is the equivalent pointer expression for refereing to same element a[i][j][k][l]?
		- A[i] can be written as *(a+i) so each can be written in a chained way.
		- A[i][j] = *(*(a+i)+j)
		- A[i][j][k][l] = *(*(*(*(a+i)+j)+k)+l);
	63. Which bitwise operation is suitable for checking bits on or off?
		- Bitwise AND (&) used to check if a bit is set or not.
		- Check the 5th bit
		- Bit = (byte >> 4) & 0x01;
	64. When should we use register modifier?
		- Used when a variable is expected to be heavily used and keeping it in the CPU's register will make a faster access.
	65. Why doesn't this work?
		- Char str[] = "Hello";
		- Strcat(str, '!');
		- '!' is a char, strcat is only for strings (i.e. char[]). "!" needs to be used instead.
	66. Predict the output or error for the following?
		- Int const *p = 5;
		- Printf("%d", ++(*p));
		- Since the int is const it will error out
	67. Guess the output?
		- Unsigned int a = 2;
		- Int b = -10;
		- (a + b > 0) ? Puts("greater than 0"): puts("less than 1");
		- Will output greater than 0 as a comes first so it's type dominates
	68. Write code fragment to clear bit 3 or integer 
		- Set bit
			§ A |= (0x1 << 3);
		- Clear bit
			§ A &= ~(0x1 << 3);
	69. What is wrong with this code?
		- Int square(volatile int *p) 
			§ Return *p * *p;
		- Tries to square the value however we might have what is at p change in each instance.
		- Should copy a local copy to square.
	70. Is this correct?
		- Int I = 2, j = 3, res;
		- Res = i+++j;
		- It will add the two values and then increment I by one afterward.


Memory Segments:
	- Text is the program code
	- Data is divided into initialized and uninitialized
		- Initialized is static and global variables with non zero values
		- Uninitialized is static and global variables with zero or NULL. Known as BSS - Block started by symbol
	- Heap and Stack are referring to RAM
		- Heap is used for dynamic memory no enforced pattern to allocation and deallocation. Grows down from high addresses.
		- Stack used for static memory allocation. In LIFO order. When a function is called a block is reserved at the top to keep data. When function returns the block is freed.


	
		
	

