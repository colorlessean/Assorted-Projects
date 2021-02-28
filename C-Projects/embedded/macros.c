#include <stdio.h>

// some macro declarations
// a const type
// a function type
// a list type

#define PI 3.14159265
#define circleArea(r) (PI*(r)*(r))
#define NUMBERS_LENGTH 3

void print(int *x) {
	for (int i = 0; i < NUMBERS_LENGTH; i++){
		printf("%d", *x);
		// increment pointer over by 1 postion
		x++;
	} 
	printf("\n");
}

int main() {
	int x[] = {1,2,3};
	print(x);
	return 0;
}
