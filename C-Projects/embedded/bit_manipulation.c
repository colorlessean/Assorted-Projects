#include <stdio.h>
#include <stdint.h>

int main() {
	uint8_t bits = 0;
	printf("example in hex: %x\n", bits);
	
	// first is to declare a bit mask
	// all 1s
	uint8_t bit_mask = 0xff;

	// to set bits we will use the |= operation
	bits |= bit_mask;
	printf("set all bits with mask: %x\n", bits);

	// next toggle the 7th and 2nd bits (LSB is 1st)
	bit_mask = 1 << 7;
	bit_mask += 1 << 2;
	bits ^= bit_mask;
	printf("toggle bits: %x\n", bits);

	// next we will clear all the bits
	bit_mask = 0xff;
	bits &= ~bit_mask;
	printf("reset all bits: %x\n", bits);

	return 0;
}
