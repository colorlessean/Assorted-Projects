#include <stdio.h>

// lets do a 3 state state machine
// kind of a thread simulation type of deal

// there will be an ready state a running state and waiting state
// special catch all state for erroneous behavior
#define STATE_ANY -1
#define STATE_READY 0
#define STATE_RUNNING 1
#define STATE_WAITING 2
#define STATE_TERMINATE 3

// events will be: blocked, unblocked, intialized
// special catch all event
#define EVENT_ANY 509
#define EVENT_TERMINATE 510
#define EVENT_BLOCK 511
#define EVENT_UNBLOCK 512

typedef struct {
	int state;
	int event;
	int (*function)(void);
} transition_t;


// declare the state transition function
// all have no input and return an int of the next state to go to
int unblock(void) {
	return STATE_WAITING;
}
int block(void) {
	return STATE_READY;
}
int init(void) {
	return STATE_READY;
}
int run(void) {
	return STATE_RUNNING;
}
int Error(void) {
	printf("AHH something broke!\n");
}

int GetNextEvent() {
	return 511;
}

#define TRANS_COUNT 7

// list of transitions that can occur
transition_t trans[] = {
	{STATE_READY, EVENT_UNBLOCK, &run},
	{STATE_READY, EVENT_BLOCK, &block},
	{STATE_RUNNING, EVENT_UNBLOCK, &run},
	{STATE_RUNNING, EVENT_BLOCK, &block},
	{STATE_WAITING, EVENT_UNBLOCK, &unblock},
	{STATE_WAITING, EVENT_BLOCK, &block},
	{STATE_ANY, EVENT_ANY, &Error}
};

int main() {
	// this is where the main loop of the state machine goes
	int state = STATE_READY;
	while (state != STATE_TERMINATE) {
		int event = GetNextEvent();
		for (int i = 0; i < TRANS_COUNT; i++) {
			if ((state == trans[i].state) || (STATE_ANY == trans[i].state)) {
				if((event == trans[i].event) || (EVENT_ANY == trans[i].event)) {
					state = (trans[i].function)();
					break;
				}
			}
		}
	}
	return 0;
}
