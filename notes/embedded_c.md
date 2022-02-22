Creating a Macro
	- Part of the Pre-processor
	- Fragment of code that has been given a name
	- i.e. the "# …. " stuff
	- Macros using #define
	- Ex: #define c 299792458 // speed of light
	- Can define lists as well like 
		○ #define NUMBERS 1,2,3
		○ Int x[] = { NUMBERS };
	- Can also make function like macros
	- #define circleArea(r) (3.1415*(r)*(r))
	- We call in the normal code like any other function
	- We can also #define PI 3.1415 and put in instead of the value in the function one as long as we declare PI first.
	- We also have the other shit like: Ifdef, if, defined, else and elseif, undef
		○ Ifdef defined, define and undef for macros
		○ If elseif and else used for expressions
		○ They can all be used with macros and expressions
	- Predifined macros are as follows:
		○ __DATE__ which gives a string of todays date
		○ __FILE__ string containing filename
		○ __LINE__ integer representing current line number
		○ __STDC__ if ANSI C will yield non-zero int
		○ __TIME__ string with the current time
	- Macros can take arguments but the functions must still be oneliners similar to lambdas
		○ #define min(X,Y) ((X) < (Y) ? (X) : (Y))
		○ We can see that this isn't really that type safe however
	- Can stringify a macro by calling #MACRO useful for printf statements
		○ Fprintf(stderr, "Warning" #EXP "\n");
	- Can concatinate tokens using the ## token pasting / token concatentation operator
		○ #define COMMAND(NAME) { #NAME, NAME ## _command }
		○ Struct command commands[] = { COMMAND (quit), COMMAND(help), … };
	- Can take in a number of arguments just a function can
		○ #define eprintf(…) fprintf (stderr, __VA_ARGS__)
		○ … sequence of tokens will replace __VA_ARGS__ 
		○ Can rename __VA_ARGS__ by simply putting a name infront of the …
		○ Can also have some named and rest variable
		○ #define eprint(format, …) fprintf(stderr, format, __VA_ARGS__)
	- https://gcc.gnu.org/onlinedocs/cpp/Macros.html


Bit Manipulation:
	- Rules for bit manipulation
	- NOT (~x), AND (x&y), OR (x|y), XOR (x^y), Left Shift (x << n), Right Shift (x >> n)
		○ Left makes bigger, right makes smaller by multiplying by 2^n bits
	- Set a bit with |= mask
	- Clear a bit with &= ~mask
	- Toggle a bit with ^= mask


Debugging:
	- Using GDB
		1. Compile the program with option -g
			i. Gcc -g program.c
		2. Launch gdb
			i. Gdb a.out
		3. Setup a breakpoint inside C program
			i. Break line_number
			ii. Break [file_name}:line_number
			iii. Break[file_name]:func_name
		4. Execute the C program in gdb
			i. Run [args]
		5. Use other gdb commands to debug
		6. Print variable values inside the gdb debugger
			i. Print {variable}
		7. Continue, step in, step over gdb commands
			i. C or continue: go to the next breakpoint hit
			ii. N or next: execute next line as single instruction
			iii. S of step : goes into the function execute it line by line
		8. Some command shortcuts
			i. L: list
			ii. P: print
			iii. C: continue
			iv. S: step
			v. ENTER: execute previously used command
			vi. 1 command: use gdb command 1 or list to print the source code in the debug mode
			vii. Use 1 line-number to view a specific line number or 1 function to view a function
			viii. Bt: backtrack print backtrace of all stack frames or innermost COUNT frames
			ix. Help: view help menu and help TOPICNAME for specifics
			x. Quit: actually exit from the debugger


Memory Dump:
	- Also called a core dump
	- Full contents of the memory when a crash occurred
	- Can also include CPU state such as registers
	- Can be used to debug
	- Allows offline debugging of a system when it crashed
	- Can be used to capture data freed duing dynamic memory allocation
	- Debugger can use a symbol table to help interpret the dump
	- GNU objdump can be a useful tool for the analysis
	- Normally the standard executable image-format
	- Can be loaded into GDB and used to debug the fault that actually occurred
		○ Gbd - start gdb with no debugging
		○ Gdb program - begin debugging a program
		○ Gbd program core - debug coredump core produced by program
		1. Find the directory where the core dump file is located
		2. Use ls -ltr command to find the latest generated corefile
		3. Load the corefile using: gbd path_to_binary path_of_core_file
		4. Gather information using the bt command
		5. Print variables using p variable_name
		6. For detailed backtrace use bt full
			i. Can also do bt n to go back by a given number of frames on stack
		7. Use frame frame-number to get the desired frame number
		8. Use up n and down n commands to select frames n frames up and n frames down


State Machine: 
	- Use function pointers
	- Need 2 arrays
		○ One for state function pointers
		○ Another for state transition rules
	- Every state function returns the code
	- Look up state transition table by state
	- Return code to find the next state
	- Execute

	- struct array and a loop
	- Structure consists of a state, event (for look-up) and function that returns new state
		Typedef struct {
			Int state;
			Int event;
			Int (*fn)(void);
		} transition_t; 
	- Can define all the states and events with simple define statements
		#define ST_ANY -1
		#define ST_INIT 0
		…
		#define EV_ANY
		#define EV_KEYPRESS
	- Could also use an enum
	- Then declare all the functions that are called by the transitions
		Static int GotKey(void) {…..};
		Static int FsmError(void) {…};
		○ Note these are all writen to take no variables and return the enw state for the state machine as they all return ints
	- Then we have a transition array that defines all the possible transitions and the functions that get called for those transitions (including the catch-all last one)
		Transition_t trans[] = {
			{ST_INIT, EV_KEYPRESS, &GotKey},
			…
			{ST_ANY, EV_ANY, &FsmError}
		};
		○ I.e. if your in the first state and you receive the second event you make a call to the third function.
	- Then the actual FSM becomes a relatively simple loop:
		State = ST_INIT;
		While (state != ST_TERM) {
			Event = GetNextEvent();
			For (I = 0; I < TRANS_COUNT; i++) {
				If ((state == trans[i].state) || (ST_ANY == trans[i].state)) {
					If ((event == trans[i].event) || (EV_ANY == trans[i].event)) {
						State  = (trans[i].fn)();
						Break;
					}
				}
			}
			
			
Unit Testing:
	- The purpose of a unit test in software engineering is to verify the behavior of a relatively small piece of software, independently from other parts. Unit tests are narrow in scope, and allow us to cover all cases, ensuring that every single part works correctly.
	- Also defined as a mthod of testing software where individual software components are isolated and tested for correctness. Ideally cover most of if not all of the code paths, argument bounds and failure cases of the SW under test
	- Can be done with or without a frame-work
	- Generally a unit test consists of:
		○ Setup and teardown functions which run before and after each test
		○ Individual tests that test logical components of paths of a module
		○ Many checks such as LONGS_EQUAL and STRCMP_EQUAL to compare int and string values
	- To make sure unit tests are funning in isolation sometimes functions and modules are not purely functional so call other shit.
	- To remedy this you can use alternative implementations of called upon modules
		○ Fakes: working implementation but will have dependencies substituted for test env.
			§ Commonly used where impractical to use real impl
		○ Stubs: trivial impl that returns canned values generally always returns valid or invalid values
			§ Used to solve linker not knowing where functions are
			§ Should only have a return statement that returns true, false, 0 or NULL
		○ Mocks: impl that is controlled by the unit test. Can be pre-programmed with return values, check values of arguments and help verify funcs are called.
			§ Useful for decalring each and every return value of a function
			§ Using many mocks is an easy way to test every single code path of the module under test as you can force any function to return error codes, NULL vals and invalid pointers
	- A framework less example which is probably most useful is:
		#include <assert.h>
		Int my_sum(int a, int b) {
			Return a + b;
		}
		Int main (int argc, char *argv[]) {
			Assert(2 == my_sum(1,1));
			Assert(-2 == my_sum(-1, -1));
			Assert(0 == my_sum(0,0));
			…
			Return(0);
		}
		
	
	
Interpolation Functions:
	
	Int main() {
		Float x0,y0,x1,y1,xp,yp;
		Clrscr();
		Printf("Enter first point then second point then interpolation point x val\n");
		Scanf("%f%f", &x0, &y0);
		Scanf("%f%f", &x1, &y1);
		Scanf("%f", &xp);
		Yp = y0 + ((y1-y0)/(x1-x0)) * (xp - x0) ;
		Printf("Interpolated point at %0.3f is %0.3f", xp, yp);
		Return 0;
	}

----------------------------------------------------------------------------------------------------------------------------------------------------
Other Stuff:

Function Pointers:
	- Can be used for callback functions
	- Pass functions as arguments
	- Contain "member functions" in a C struct
	Void (*foo)(int);
		○ Foo is a pointer to a function taking one arg, an int, that returns void
	Void *(*foo)(int *);
		○ Key is to read inside out
		○ * foo is a function that returns void * and take in an int *
		○ Foo is a reference to such a function
	- How to initalize a function pointer
		Void my_int_func(int x) { … }
		Int main() {
			Void (*foo)(int);
			Foo = &my_int_func;
			…
		}
	- Function pointer with input and output types
		int my_func(double x, void * y) { ... }
		int main() {
			int (*my_func)(double, void *);
			func = &my_func
		}
	- Can be used in either of the two ways
		Foo( args…);
		(*foo)(args…);
	- In the wild we can see the most popular is in the qsort defintion
		void qsort(void *base, size_t nmemb, size_t size, int(*compar)(const void *, const void *));
		○ Compar is a function you must pass to the function
		○ You can just pass any function without any & stuff tho from your main functions
	
C Sockets:
	To create a new socket
		1. Call socket()
		2. Call bind()
		3. Call listen()
		4. can then call a number of
			send()
			recv()
	To accept new connections the socket must
		1. accept()
	To connect to an existing socket
		1. Call socket()
		2. Call connect()
		3. Call a number of
			send()
			recv()

